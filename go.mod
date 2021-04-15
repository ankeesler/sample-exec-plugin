module github.com/ankeesler/sample-exec-plugin

go 1.15

require (
	k8s.io/api v0.21.0 // indirect
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
)

replace (
	k8s.io/api => /Users/akeesler/workspace/src/k8s.io/kubernetes/staging/src/k8s.io/api
	k8s.io/apimachinery => /Users/akeesler/workspace/src/k8s.io/kubernetes/staging/src/k8s.io/apimachinery
	k8s.io/client-go => /Users/akeesler/workspace/src/k8s.io/kubernetes/staging/src/k8s.io/client-go
)
