#person: {
	name:  string
	age:   int & >=0
	human: true
    ...
}

ichigo: #person & {
	name: "Ichigo Hoshimiya"
	age:  22
    foo: "Bar"
}
