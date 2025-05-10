#foo: 10
#bar: 3

result: "This is the result of foo - bar: \(#foo-#bar)"

foo: {
	bar: {
		longText: """
        Hello world
        line 2
        line 3
        interoperable text: \(result)
        """
	}
}

a:       "foo"
b:       "bar"
(a + b): "foobar"

s: X={
	"\(a)_and_\(b)": "foobar"

	// Valid references using a selector and
	// an index expression.
	FooAndBar: s.foo_and_bar
	FooAndBar: X["foo_and_bar"]
	FooAndBar: s["\(a)_and_\(b)"]

	// Invalid reference because the
	// indentifer is not in scope.
	//FooAndBar: foo_and_bar
}

// Valid reference using an index expression.
FooAndBar: s["foo_and_bar"]

d1: #"\U0001F60E"#
d2: "\U0001F60E"
d3: #"""
asdf
fdsa
\U0001F60E
"""#
