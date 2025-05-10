aString: "A string"

// Multiple representations of the same underlying
// bytes, which therefore unify succesfully.
Bytes: 'A string'
Bytes: '\(aString)'
Bytes: '\x41\x20\x73\x74\x72\x69\x6e\x67'
Bytes: '\101\040\163\164\162\151\156\147'

stringBytes: "\(Bytes)"
