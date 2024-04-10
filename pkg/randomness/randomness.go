package randomness

import (
	math_rand "math/rand/v2"
)

// lengths
const (
	lengthDomain = 63
)

// character sets
const (
	charSetDomain = "0123456789abcdefghijklmnopqrstuvwxyz" // case insensitive; exclude `-` to avoid any unexpected anomalies
)

// Insecure Randoms

// GenerateRandomComDomain creates a random (sub)domain in the format
// <random-string-1>.<random-string-2>.com.
func GenerateRandomComDomain() string {
	fqdn := ""

	for i := 0; i < 2; i++ {
		bytes := make([]byte, lengthDomain)
		for i := range bytes {
			bytes[i] = charSetDomain[math_rand.IntN(len(charSetDomain))]
		}

		fqdn += string(bytes) + "."
	}

	fqdn += "com."

	return fqdn
}
