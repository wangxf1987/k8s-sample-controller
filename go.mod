module k8s.io/k8s-sample-controller

go 1.13

require (
	k8s.io/api v0.18.5
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v0.18.5
	k8s.io/code-generator v0.18.5
	k8s.io/klog v1.0.0
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.18.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.5
	k8s.io/client-go => k8s.io/client-go v0.18.5
	k8s.io/code-generator => k8s.io/code-generator v0.18.5
)
