module github.com/ankeesler/sample-exec-plugin

go 1.15

require (
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
)

replace (
	k8s.io/api => /Users/akeesler/workspace/src/k8s.io/kubernetes/staging/src/k8s.io/api
	k8s.io/apimachinery => /Users/akeesler/workspace/src/k8s.io/kubernetes/staging/src/k8s.io/apimachinery
	k8s.io/client-go => /Users/akeesler/workspace/src/k8s.io/kubernetes/staging/src/k8s.io/client-go
)