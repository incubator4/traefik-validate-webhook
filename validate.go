package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// validate deployments and services
func (whsvr *WebhookServer) validate(ar *admissionv1.AdmissionReview) *admissionv1.AdmissionResponse {
	req := ar.Request
	var (
		objectMeta                      *metav1.ObjectMeta
		resourceNamespace, resourceName string
	)

	var routes []v1alpha1.Route

	glog.Infof("AdmissionReview for Kind=%v, Namespace=%v Name=%v (%v) UID=%v patchOperation=%v UserInfo=%v",
		req.Kind, req.Namespace, req.Name, resourceName, req.UID, req.Operation, req.UserInfo)

	switch req.Kind.Kind {
	case "IngressRoute":
		var ingressRoute v1alpha1.IngressRoute
		if err := json.Unmarshal(req.Object.Raw, &ingressRoute); err != nil {
			glog.Errorf("Could not unmarshal raw object: %v", err)
			return &admissionv1.AdmissionResponse{
				Result: &metav1.Status{
					Message: err.Error(),
				},
			}
		}
		resourceName, resourceNamespace, objectMeta = ingressRoute.Name, ingressRoute.Namespace, &ingressRoute.ObjectMeta
		routes = ingressRoute.Spec.Routes
	}

	if !validationRequired(ignoredNamespaces, objectMeta) {
		glog.Infof("Skipping validation for %s/%s due to policy check", resourceNamespace, resourceName)
		return &admissionv1.AdmissionResponse{
			Allowed: true,
		}
	}

	return validationRoutes(routes)
}

func validationRequired(ignoredList []string, metadata *metav1.ObjectMeta) bool {
	required := admissionRequired(ignoredList, admissionWebhookAnnotationValidateKey, metadata)
	glog.Infof("Validation policy for %v/%v: required:%v", metadata.Namespace, metadata.Name, required)
	return required
}

func validationRoutes(routes []v1alpha1.Route) *admissionv1.AdmissionResponse {
	ruleMap, err := ListRules()
	if err != nil {
		glog.Errorf("Cannot list rules error: %s", err.Error())
		return &admissionv1.AdmissionResponse{Allowed: true}
	}

	for _, route := range routes {
		for _, r := range SplitMatchPath(route.Match) {
			if ruleMap[r.ToString()] != nil {
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
