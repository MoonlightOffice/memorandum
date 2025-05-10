package main

import (
	"k8s.example/svc"
	"k8s.example/deploy"
	"encoding/yaml"
	"tool/cli"
)



resources: [
    for resource in svc.resources {resource}
    for resource in deploy.resources {resource}
]

command: dump: {
	task: print: cli.Print & {
		text: yaml.MarshalStream(resources)
	}
}

