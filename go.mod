module github.com/hwameistor/hwameistor-ui

go 1.18

replace (
	k8s.io/api => k8s.io/api v0.23.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.23.6
	k8s.io/apiserver => k8s.io/apiserver v0.23.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.6
	k8s.io/client-go => k8s.io/client-go v0.23.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.6
	k8s.io/code-generator => k8s.io/code-generator v0.23.6
	k8s.io/component-base => k8s.io/component-base v0.23.6
	k8s.io/component-helpers => k8s.io/component-helpers v0.23.6
	k8s.io/controller-manager => k8s.io/controller-manager v0.23.6
	k8s.io/cri-api => k8s.io/cri-api v0.23.6
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.6
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.6
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.6
	k8s.io/kubectl => k8s.io/kubectl v0.23.6
	k8s.io/kubelet => k8s.io/kubelet v0.23.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.6
	k8s.io/metrics => k8s.io/metrics v0.23.6
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.6
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.6
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/handlers v1.4.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/longhorn/backing-image-manager v0.0.0-20220609065820-a08f7f47442f // indirect
	github.com/longhorn/backupstore v0.0.0-20220913112826-5f5c95274f2a // indirect
	github.com/longhorn/go-iscsi-helper v0.0.0-20220927074943-051bf960608b // indirect
	github.com/longhorn/longhorn-instance-manager v0.0.0-20220929053134-65135bc070cd // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/rancher/go-rancher v0.1.1-0.20220412083059-ff12399dd57b
	github.com/robfig/cron v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.24.2 // indirect
	k8s.io/apiextensions-apiserver v0.24.0 // indirect
	k8s.io/apimachinery v0.24.2 // indirect
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kubernetes v1.23.6 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/controller-runtime v0.10.1
	sigs.k8s.io/yaml v1.3.0 // indirect
)

require (
	github.com/hwameistor/hwameistor v0.3.7-rc.8
	github.com/longhorn/longhorn-manager v1.3.2
)

require (
	github.com/RoaringBitmap/roaring v0.4.18 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/c9s/goprocinfo v0.0.0-20190309065803-0b2ad9ac246b // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/container-storage-interface/spec v1.5.0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/glycerine/go-unsnap-stream v0.0.0-20181221182339-f9677308dec2 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/honestbee/jobq v1.0.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/longhorn/longhorn-engine v1.3.2-0.20220929032851-7aac8ae9c8b4 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mschoch/smat v0.0.0-20160514031455-90eadee771ae // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/runc v1.0.2 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tinylib/msgp v1.1.1-0.20190612170807-0573788bc2a8 // indirect
	github.com/willf/bitset v1.1.10 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/apiserver v0.23.6 // indirect
	k8s.io/component-base v0.23.6 // indirect
	k8s.io/component-helpers v0.23.6 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65 // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)
