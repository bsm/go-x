package hashx

// Copied from https://github.com/pierrec/xxHash
// Licensed under BSD 3-clause.

const (
	xxPrime32_1 = 2654435761
	xxPrime32_2 = 2246822519
	xxPrime32_3 = 3266489917
	xxPrime32_4 = 668265263
	xxPrime32_5 = 374761393
)

// XX32 returns the 32bit xxHash value.
func XX32(input []byte, seed uint32) uint32 {
	n := len(input)
	if n == 0 {
		return 0
	}

	h32 := uint32(n)
	if n < 16 {
		h32 += seed + xxPrime32_5
	} else {
		v1 := seed + xxPrime32_1 + xxPrime32_2
		v2 := seed + xxPrime32_2
		v3 := seed
		v4 := seed - xxPrime32_1
		p := 0
		for n := n - 16; p <= n; p += 16 {
			sub := input[p:][:16] //BCE hint for compiler
			v1 = rol32_13(v1+u32(sub[:])*xxPrime32_2) * xxPrime32_1
			v2 = rol32_13(v2+u32(sub[4:])*xxPrime32_2) * xxPrime32_1
			v3 = rol32_13(v3+u32(sub[8:])*xxPrime32_2) * xxPrime32_1
			v4 = rol32_13(v4+u32(sub[12:])*xxPrime32_2) * xxPrime32_1
		}
		input = input[p:]
		n -= p
		h32 += rol32_1(v1) + rol32_7(v2) + rol32_12(v3) + rol32_18(v4)
	}

	p := 0
	for n := n - 4; p <= n; p += 4 {
		h32 += u32(input[p:p+4]) * xxPrime32_3
		h32 = rol32_17(h32) * xxPrime32_4
	}
	for p < n {
		h32 += uint32(input[p]) * xxPrime32_5
		h32 = rol32_11(h32) * xxPrime32_1
		p++
	}

	h32 ^= h32 >> 15
	h32 *= xxPrime32_2
	h32 ^= h32 >> 13
	h32 *= xxPrime32_3
	h32 ^= h32 >> 16

	return h32
}

const (
	xxPrime64_1 = 11400714785074694791
	xxPrime64_2 = 14029467366897019727
	xxPrime64_3 = 1609587929392839161
	xxPrime64_4 = 9650029242287828579
	xxPrime64_5 = 2870177450012600261
)

// XX64 returns the 64bit xxHash value.
func XX64(data []byte, seed uint64) uint64 {
	n := len(data)
	if n == 0 {
		return 0
	}

	var h64 uint64
	if n >= 32 {
		v1 := seed + xxPrime64_1 + xxPrime64_2
		v2 := seed + xxPrime64_2
		v3 := seed
		v4 := seed - xxPrime64_1
		p := 0
		for n := n - 32; p <= n; p += 32 {
			sub := data[p:][:32] //BCE hint for compiler
			v1 = rol64_31(v1+u64(sub[:])*xxPrime64_2) * xxPrime64_1
			v2 = rol64_31(v2+u64(sub[8:])*xxPrime64_2) * xxPrime64_1
			v3 = rol64_31(v3+u64(sub[16:])*xxPrime64_2) * xxPrime64_1
			v4 = rol64_31(v4+u64(sub[24:])*xxPrime64_2) * xxPrime64_1
		}

		h64 = rol64_1(v1) + rol64_7(v2) + rol64_12(v3) + rol64_18(v4)

		v1 *= xxPrime64_2
		v2 *= xxPrime64_2
		v3 *= xxPrime64_2
		v4 *= xxPrime64_2

		h64 = (h64^(rol64_31(v1)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64_31(v2)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64_31(v3)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64_31(v4)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4

		h64 += uint64(n)

		data = data[p:]
		n -= p
	} else {
		h64 = seed + xxPrime64_5 + uint64(n)
	}

	p := 0
	for n := n - 8; p <= n; p += 8 {
		sub := data[p : p+8]
		h64 ^= rol64_31(u64(sub)*xxPrime64_2) * xxPrime64_1
		h64 = rol64_27(h64)*xxPrime64_1 + xxPrime64_4
	}
	if p+4 <= n {
		sub := data[p : p+4]
		h64 ^= uint64(u32(sub)) * xxPrime64_1
		h64 = rol64_23(h64)*xxPrime64_2 + xxPrime64_3
		p += 4
	}
	for ; p < n; p++ {
		h64 ^= uint64(data[p]) * xxPrime64_5
		h64 = rol64_11(h64) * xxPrime64_1
	}

	h64 ^= h64 >> 33
	h64 *= xxPrime64_2
	h64 ^= h64 >> 29
	h64 *= xxPrime64_3
	h64 ^= h64 >> 32

	return h64
}
