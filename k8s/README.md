# Kubernetes Cheat Sheet & Templates

## Basic commands

Read docs
```shell
kubectl api-resources
kubectl explain pod
kubectl explain pod.spec.containers
```

Get resources
```shell
kubectl get pod

kubectl get pvc
kubectl get PersistentVolumeClaim

kubectl get svc
kubectl get Service

kubectl get deploy
kubectl get Deployment

kubectl get sts
kubectl get StatefulSet

kubectl get ns
kubectl get Namespace

kubectl get pod -n <namespace>
kubectl get pod --namespace <namespace>

kubectl get all

kubectl describe pod
kubectl describe pod <pod-name>
```

Delete resources
```shell
kubectl delete pod <pod-name>
kubectl delete --all pod
kubectl delete pvc <pvc-name>
# Delete by label
kubectl delete pvc -l key1=value1,key2=value2
```

Read logs
```shell
kubectl logs -f <pod-name>
```

Execute arbitrary commands in container
```shell
kubectl exec <pod-name> -- ls
kubectl exec -ti <pod-name> -- sh
kubectl exec -ti <pod-name> -c <container-name> -- sh
```

Work with yaml files
```shell
kubectl apply -f example.yaml
kubectl apply -f example/ -R
kubectl delete -f example.yaml

# Specify namespace
kubectl apply -f example.yaml -n <namespace>
kubectl delete -f example.yaml -n <namespace>
```

## Kubernetes internal DNS

Pod's DNS pattern is as follows:
```shell
# Pods with ordinary service
<service-name>.<namespace>.svc.cluster.local

# Pods with StatefulSets
<pod-name>.<headless-service-name>.<namespace>.svc.cluster.local
```

For example, if a pod is trying to access another pod named "server" with the service name "lb" and located in the "default" namespace, it can reach the pod using the following domain name:
```shell
# Pods with ordinary service
lb.default.svc.cluster.local

# Pods with StatefulSets
server.lb.default.svc.cluster.local
```

If both pods are in the same namespace, the DNS can be shortened to:

```shell
# Pods with ordinary service
lb

# Pods with StatefulSets
server.lb
```