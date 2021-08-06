# GoRestApi

## Setting up Kubernetes locally with kind
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/) is the defacto standard for running K8S cluster locally on your machine for testing purposes.

Install Kind with brew.

```shell
brew install kind
```

After the successful installation, create a K8S cluster.

```shell
kind create cluster
```

That creates a cluster with the name `kind` in the local machine. Test cluster creation with the following command:

```shell
kubectl cluster-info --context kind-kind
```