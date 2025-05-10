n1: L= {
	n2e1: "n2e1-val"
	n2e2: {
		n3e1: "n3e1-val"
	}
	n2e3: {
		n3e1: n2e2.n3e1
		n3e2: L.n2e2
	}
}
