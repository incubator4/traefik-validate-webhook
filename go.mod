module github.com/incubator4/traefik-validate-webhook

go 1.19

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.15.0
	github.com/traefik/traefik/v2 v2.9.10
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/kubernetes v1.22.1
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-acme/lego/v4 v4.10.2 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/go-logr/logr v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/rogpeppe/go-internal v1.6.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/traefik/paerser v0.2.0 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/tools v0.2.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.21.3 // indirect
	k8s.io/apiserver v0.22.1 // indirect
	k8s.io/component-base v0.22.1 // indirect
	k8s.io/klog/v2 v2.10.0 // indirect
	k8s.io/utils v0.0.0-20210820185131-d34e5cb4466e // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.2 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/abbot/go-http-auth => github.com/containous/go-http-auth v0.4.1-0.20200324110947-a37a7636d23e
	github.com/go-check/check => github.com/containous/check v0.0.0-20170915194414-ca0bf163426a
	github.com/mailgun/minheap => github.com/containous/minheap v0.0.0-20190809180810-6e71eb837595
	k8s.io/api => k8s.io/api v0.22.1
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.1
	k8s.io/apiserver => k8s.io/apiserver v0.22.1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.22.1
	k8s.io/client-go => k8s.io/client-go v0.22.1
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.22.1
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.22.1
	k8s.io/code-generator => k8s.io/code-generator v0.22.1
	k8s.io/component-base => k8s.io/component-base v0.22.1
	k8s.io/component-helpers => k8s.io/component-helpers v0.22.1
	k8s.io/controller-manager => k8s.io/controller-manager v0.22.1
	k8s.io/cri-api => k8s.io/cri-api v0.22.1
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.22.1
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.22.1
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.22.1
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.22.1
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.22.1
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.22.1
	k8s.io/kubectl => k8s.io/kubectl v0.22.1
	k8s.io/kubelet => k8s.io/kubelet v0.22.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.22.1
	k8s.io/metrics => k8s.io/metrics v0.22.1
	k8s.io/mount-utils => k8s.io/mount-utils v0.22.1
	k8s.io/node-api => k8s.io/node-api v0.22.1
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.22.1
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.22.1
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.22.1
	k8s.io/sample-controller => k8s.io/sample-controller v0.22.1
)
