import (
	"math"
)

// Simple labels don't need to be quoted.

one:        1
two:        2
piePlusOne: math.Pi + one

"quoted field names": {
	"two-and-a-half":    2.5
	"three point three": 3.3
	"four^four":         math.Pow(4, 4)
}

aList: [1, 2, 3]

metadata: {
	name: string
	labels: [string]
}
