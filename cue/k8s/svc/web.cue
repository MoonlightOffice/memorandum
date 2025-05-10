package svc

web: { 
    apiVersion: "v1"
    kind: "Service"
    metadata: {
        name: "web"
        labels: app: "web"
    }
    spec: {
        type: "ClusterIP"
        selector: app: "web"
        ports: [
            {
                port: 8080
                targetPort: 8080
            }
        ]
    }
}

