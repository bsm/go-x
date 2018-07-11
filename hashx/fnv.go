package hashx

const (
	fnvOffset32 uint32 = 2166136261
	fnvPrime32  uint32 = 16777619

	fnvOffset64 uint64 = 14695981039346656037
	fnvPrime64  uint64 = 1099511628211
)

// Fnv64a calculates 64-bit FNV-1a hash.
// It returns 0 on empty input.
func Fnv64a(b []byte) uint64 {
	if len(b) == 0 {
		return 0
	}

	h := fnvOffset64
	for _, c := range b {
		h ^= uint64(c)
		h *= fnvPrime64
	}
	return h
}

// Fnv64aString calculates 64-bit FNV-1a hash.
// It returns 0 on empty input.
func Fnv64aString(s string) uint64 {
	if len(s) == 0 {
		return 0
	}

	h := fnvOffset64
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= fnvPrime64
	}
	return h
}

// Fnv32a calculates 32-bit FNV-1a hash.
// It returns 0 on empty input (nil or empty byte slice).
func Fnv32a(b []byte) uint32 {
	if len(b) == 0 {
		return 0
	}

	h := fnvOffset32
	for _, c := range b {
		h ^= uint32(c)
		h *= fnvPrime32
	}
	return h
}

// Fnv32aString calculates 32-bit FNV-1a hash.
// It returns 0 on empty input (nil or empty strings).
func Fnv32aString(s string) uint32 {
	if len(s) == 0 {
		return 0
	}

	h := fnvOffset32
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= fnvPrime32
	}
	return h
}
