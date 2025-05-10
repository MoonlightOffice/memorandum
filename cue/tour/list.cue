#n: [1, 2, 3, 4, 5, 6, 7, 8, 9]
#s: ["a", "b", "c"]

// Large numbers
a: [
	for x in #n
	let xx = x * x
	if xx > 50 {xx},
]

// Even numbers
b: [for x in #n if mod(x, 2) == 0 {x}]

// Cartesian product
c: [
	for x in #n
	for y in #s
	if x <= 3 {
		"\(y)-\(x)"
	},
]
