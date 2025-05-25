## Install Ingress controller

```shell
# Add helm repository to local if necessary
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

# Install or upgrade ingress controller in Kubernetes cluster
helm upgrade --install -n ingress-controller --create-namespace --values values.yaml ingress-nginx ingress-nginx/ingress-nginx

# Uninstall ingress controller in Kubernetes cluster
helm -n ingress-controller uninstall ingress-nginx
```

`values.yaml` allows you to specify the number of replicas of both ingress controller pods and cloud load balancer nodes.