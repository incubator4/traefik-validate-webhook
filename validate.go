package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"
)

// validate deployments and services
func (whsvr *WebhookServer) validate(ar *admissionv1.AdmissionReview) *admissionv1.AdmissionResponse {
	req := ar.Request
	var (
		objectMeta                      *metav1.ObjectMeta
		resourceNamespace, resourceName string
		ingressRoute                    v1alpha1.IngressRoute
	)

	glog.Infof("AdmissionReview for Kind=%v, Namespace=%v Name=%v (%v) UID=%v patchOperation=%v UserInfo=%v",
		req.Kind, req.Namespace, req.Name, resourceName, req.UID, req.Operation, req.UserInfo)

	switch req.Kind.Kind {
	case "IngressRoute":

		if err := json.Unmarshal(req.Object.Raw, &ingressRoute); err != nil {
			glog.Errorf("Could not unmarshal raw object: %v", err)
			return &admissionv1.AdmissionResponse{
				Result: &metav1.Status{
					Message: err.Error(),
				},
			}
		}
		resourceName, resourceNamespace, objectMeta = ingressRoute.Name, ingressRoute.Namespace, &ingressRoute.ObjectMeta
	}

	if !validationRequired(ignoredNamespaces, objectMeta) {
		glog.Infof("Skipping validation for %s/%s due to policy check", resourceNamespace, resourceName)
		return &admissionv1.AdmissionResponse{
			Allowed: true,
		}
	}

	return validationRoutes(ingressRoute)
}

func validationRequired(ignoredList []string, metadata *metav1.ObjectMeta) bool {
	required := admissionRequired(ignoredList, admissionWebhookAnnotationValidateKey, metadata)
	glog.Infof("Validation policy for %v/%v: required:%v", metadata.Namespace, metadata.Name, required)
	return required
}

func validationRoutes(ing v1alpha1.IngressRoute) *admissionv1.AdmissionResponse {
	ruleMap, err := ListRules()
	if err != nil {
		glog.Errorf("Cannot list rules error: %s", err.Error())
		return &admissionv1.AdmissionResponse{Allowed: true}
	}
	entryPoints, err := ListEntryPoints()
	if err != nil {
		glog.Errorf("Cannot list entrypoints error: %s", err.Error())
		return &admissionv1.AdmissionResponse{Allowed: true}
	}

	for _, route := range ing.Spec.Routes {
		for _, r := range SplitMatchPath(route.Match) {
			rule := ruleMap[r.ToString()]
			isOwner := validateOwner(ing, rule, entryPoints)
			glog.Infof("rule %+v rule empty %t, owner %t", rule, rule.IsEmpty(), isOwner)
			if !rule.IsEmpty() && !isOwner {
				glog.Warningf("detect duplicate route %s", r.ToString())
				return &admissionv1.AdmissionResponse{
					Allowed: false,
					Result: &metav1.Status{
						Message: fmt.Sprintf("route %s has already defined", r.ToString()),
						Reason:  metav1.StatusReasonAlreadyExists,
					},
				}
			}
		}
	}

	return &admissionv1.AdmissionResponse{Allowed: true}
}

func validateOwner(ing v1alpha1.IngressRoute, route Route, eps []EntryPoint) bool {

	var prefixBuffer = new(bytes.Buffer)
	for i, e := range eps {
		if i > 0 {
			prefixBuffer.WriteString("|")
		}
		prefixBuffer.WriteString(e.Name)
	}

	reStr := fmt.Sprintf(
		`^(%s)(?P<namespace>%s)-(?P<name>%s)-(?P<hash>[0-9a-f]+)@kubernetescrd$`,
		prefixBuffer.String(),
		ing.Namespace, ing.Name,
	)
	re := regexp.MustCompile(reStr)
	match := re.FindStringSubmatch(route.Owner)
	if len(match) > 0 {
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		if result["name"] == ing.Name && result["namespace"] == ing.Namespace {
			return true
		}
	}
	return false
}
