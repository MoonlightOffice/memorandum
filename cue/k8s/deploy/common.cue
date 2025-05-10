package deploy

#DeploymentParam: {
    name!: string
    replicas: >= 1 | *1
    ...
}

#Deployment: p=#DeploymentParam & {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata: {
        namespace: "default"
		name: p.name
		labels: app: p.name
	}
	spec: {
        replicas: p.replicas
		strategy: {
			type: "RollingUpdate"
			rollingUpdate: {
				maxUnavailable: "25%"
				maxSurge:       "25%"
			}
		}
		selector: matchLabels: {
			app: "web"
		}
		template: {
			metadata: {
				name: "web"
				labels: app: "web"
			}
			spec: {
				containers: [
					{
						name:            "web"
						image:           "caddy"
						imagePullPolicy: "IfNotPresent"
						command: ["caddy", "file-server", "--listen", "0.0.0.0:8080"]
					},
				]
				terminationGracePeriodSeconds: 1
			}
		}
	}
}