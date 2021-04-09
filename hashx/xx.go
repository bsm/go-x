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
			v1 = rol32(v1+u32(sub[:])*xxPrime32_2, 13) * xxPrime32_1
			v2 = rol32(v2+u32(sub[4:])*xxPrime32_2, 13) * xxPrime32_1
			v3 = rol32(v3+u32(sub[8:])*xxPrime32_2, 13) * xxPrime32_1
			v4 = rol32(v4+u32(sub[12:])*xxPrime32_2, 13) * xxPrime32_1
		}
		input = input[p:]
		n -= p
		h32 += rol32(v1, 1) + rol32(v2, 7) + rol32(v3, 12) + rol32(v4, 18)
	}

	p := 0
	for n := n - 4; p <= n; p += 4 {
		h32 += u32(input[p:p+4]) * xxPrime32_3
		h32 = rol32(h32, 17) * xxPrime32_4
	}
	for p < n {
		h32 += uint32(input[p]) * xxPrime32_5
		h32 = rol32(h32, 11) * xxPrime32_1
		p++
	}

	h32 ^= h32 >> 15
	h32 *= xxPrime32_2
	h32 ^= h32 >> 13
	h32 *= xxPrime32_3
	h32 ^= h32 >> 16

	return h32
}

// XX32String returns the 32bit xxHash value.
func XX32String(input string, seed uint32) uint32 {
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
			v1 = rol32(v1+u32s(sub, 0)*xxPrime32_2, 13) * xxPrime32_1
			v2 = rol32(v2+u32s(sub, 4)*xxPrime32_2, 13) * xxPrime32_1
			v3 = rol32(v3+u32s(sub, 8)*xxPrime32_2, 13) * xxPrime32_1
			v4 = rol32(v4+u32s(sub, 12)*xxPrime32_2, 13) * xxPrime32_1
		}
		input = input[p:]
		n -= p
		h32 += rol32(v1, 1) + rol32(v2, 7) + rol32(v3, 12) + rol32(v4, 18)
	}

	p := 0
	for n := n - 4; p <= n; p += 4 {
		h32 += u32s(input[p:p+4], 0) * xxPrime32_3
		h32 = rol32(h32, 17) * xxPrime32_4
	}
	for p < n {
		h32 += uint32(input[p]) * xxPrime32_5
		h32 = rol32(h32, 11) * xxPrime32_1
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
			v1 = rol64(v1+u64(sub[:])*xxPrime64_2, 31) * xxPrime64_1
			v2 = rol64(v2+u64(sub[8:])*xxPrime64_2, 31) * xxPrime64_1
			v3 = rol64(v3+u64(sub[16:])*xxPrime64_2, 31) * xxPrime64_1
			v4 = rol64(v4+u64(sub[24:])*xxPrime64_2, 31) * xxPrime64_1
		}

		h64 = rol64(v1, 1) + rol64(v2, 7) + rol64(v3, 12) + rol64(v4, 18)

		v1 *= xxPrime64_2
		v2 *= xxPrime64_2
		v3 *= xxPrime64_2
		v4 *= xxPrime64_2

		h64 = (h64^(rol64(v1, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64(v2, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64(v3, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64(v4, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4

		h64 += uint64(n)

		data = data[p:]
		n -= p
	} else {
		h64 = seed + xxPrime64_5 + uint64(n)
	}

	p := 0
	for n := n - 8; p <= n; p += 8 {
		sub := data[p : p+8]
		h64 ^= rol64(u64(sub)*xxPrime64_2, 31) * xxPrime64_1
		h64 = rol64(h64, 27)*xxPrime64_1 + xxPrime64_4
	}
	if p+4 <= n {
		sub := data[p : p+4]
		h64 ^= uint64(u32(sub)) * xxPrime64_1
		h64 = rol64(h64, 23)*xxPrime64_2 + xxPrime64_3
		p += 4
	}
	for ; p < n; p++ {
		h64 ^= uint64(data[p]) * xxPrime64_5
		h64 = rol64(h64, 11) * xxPrime64_1
	}

	h64 ^= h64 >> 33
	h64 *= xxPrime64_2
	h64 ^= h64 >> 29
	h64 *= xxPrime64_3
	h64 ^= h64 >> 32

	return h64
}

// XX64String returns the 64bit xxHash value.
func XX64String(data string, seed uint64) uint64 {
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
			v1 = rol64(v1+u64s(sub, 0)*xxPrime64_2, 31) * xxPrime64_1
			v2 = rol64(v2+u64s(sub, 8)*xxPrime64_2, 31) * xxPrime64_1
			v3 = rol64(v3+u64s(sub, 16)*xxPrime64_2, 31) * xxPrime64_1
			v4 = rol64(v4+u64s(sub, 24)*xxPrime64_2, 31) * xxPrime64_1
		}

		h64 = rol64(v1, 1) + rol64(v2, 7) + rol64(v3, 12) + rol64(v4, 18)

		v1 *= xxPrime64_2
		v2 *= xxPrime64_2
		v3 *= xxPrime64_2
		v4 *= xxPrime64_2

		h64 = (h64^(rol64(v1, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64(v2, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64(v3, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4
		h64 = (h64^(rol64(v4, 31)*xxPrime64_1))*xxPrime64_1 + xxPrime64_4

		h64 += uint64(n)

		data = data[p:]
		n -= p
	} else {
		h64 = seed + xxPrime64_5 + uint64(n)
	}

	p := 0
	for n := n - 8; p <= n; p += 8 {
		sub := data[p : p+8]
		h64 ^= rol64(u64s(sub, 0)*xxPrime64_2, 31) * xxPrime64_1
		h64 = rol64(h64, 27)*xxPrime64_1 + xxPrime64_4
	}
	if p+4 <= n {
		sub := data[p : p+4]
		h64 ^= uint64(u32s(sub, 0)) * xxPrime64_1
		h64 = rol64(h64, 23)*xxPrime64_2 + xxPrime64_3
		p += 4
	}
	for ; p < n; p++ {
		h64 ^= uint64(data[p]) * xxPrime64_5
		h64 = rol64(h64, 11) * xxPrime64_1
	}

	h64 ^= h64 >> 33
	h64 *= xxPrime64_2
	h64 ^= h64 >> 29
	h64 *= xxPrime64_3
	h64 ^= h64 >> 32

	return h64
}
