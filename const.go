package pronto

const (
	// allowedSignalDeviation is the percentage deviation allowed to consider two
	// signals still equal.
	allowedSignalDeviation = 0.05

	// minimumSignalDeviation is the smallest absolute deviation value to override
	// the allowedSignalDeviation on small numbers (10% of 10 is 1, but 10% of 10000 is 1000).
	minimumSignalDeviation = 4

	// burstPairSequenceLength is the length of burst pair in hexadecimal characters.
	burstPairSequenceLength = 4
)
