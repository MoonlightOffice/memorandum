job: {
	nginx: replicas:  2
	manager: command: "monit -I"
	policyd: _
	boltdb: replicas:   3
	postgres: replicas: 5
}

job: [Name=string]: {
	name:     Name // Name is an alias.
	command:  string | *"exec \(Name)"
	replicas: uint | *1
}

// Databases are important, so increase the
// replica minimum.
job: [#DB]: replicas: >=3
#DB: "postgres" | "mysql" | =~"db$"
