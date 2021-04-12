package hashx

// Taken from https://github.com/dgryski/go-metro
// Copyright (c) 2016 Damian Gryski
// MIT

const (
	metroK0 = 0xD6D018F5
	metroK1 = 0xA2AA033B
	metroK2 = 0x62992FC1
	metroK3 = 0x30BC5B29
)

// Metro64 returns the 64bit metro hash value.
func Metro64(buffer []byte, seed uint64) uint64 {
	ptr := buffer
	hash := (seed + metroK2) * metroK0
	if len(ptr) >= 32 {
		v0, v1, v2, v3 := hash, hash, hash, hash

		for len(ptr) >= 32 {
			v0 += u64(ptr[:8]) * metroK0
			v0 = rol64(v0, -29) + v2
			v1 += u64(ptr[8:16]) * metroK1
			v1 = rol64(v1, -29) + v3
			v2 += u64(ptr[16:24]) * metroK2
			v2 = rol64(v2, -29) + v0
			v3 += u64(ptr[24:32]) * metroK3
			v3 = rol64(v3, -29) + v1
			ptr = ptr[32:]
		}

		v2 ^= rol64(((v0+v3)*metroK0)+v1, -37) * metroK1
		v3 ^= rol64(((v1+v2)*metroK1)+v0, -37) * metroK0
		v0 ^= rol64(((v0+v2)*metroK0)+v3, -37) * metroK1
		v1 ^= rol64(((v1+v3)*metroK1)+v2, -37) * metroK0
		hash += v0 ^ v1
	}

	if len(ptr) >= 16 {
		v0 := hash + (u64(ptr[:8]) * metroK2)
		v0 = rol64(v0, -29) * metroK3
		v1 := hash + (u64(ptr[8:16]) * metroK2)
		v1 = rol64(v1, -29) * metroK3
		v0 ^= rol64(v0*metroK0, -21) + v1
		v1 ^= rol64(v1*metroK3, -21) + v0
		hash += v1
		ptr = ptr[16:]
	}

	if len(ptr) >= 8 {
		hash += u64(ptr[:8]) * metroK3
		ptr = ptr[8:]
		hash ^= rol64(hash, -55) * metroK1
	}

	if len(ptr) >= 4 {
		hash += uint64(u32(ptr[:4])) * metroK3
		hash ^= rol64(hash, -26) * metroK1
		ptr = ptr[4:]
	}

	if len(ptr) >= 2 {
		hash += uint64(u16(ptr[:2])) * metroK3
		ptr = ptr[2:]
		hash ^= rol64(hash, -48) * metroK1
	}

	if len(ptr) >= 1 {
		hash += uint64(ptr[0]) * metroK3
		hash ^= rol64(hash, -37) * metroK1
	}

	hash ^= rol64(hash, -28)
	hash *= metroK0
	hash ^= rol64(hash, -29)

	return hash
}

// Metro64String returns the 64bit metro hash value.
func Metro64String(s string, seed uint64) uint64 {
	ptr := s
	hash := (seed + metroK2) * metroK0
	if len(ptr) >= 32 {
		v0, v1, v2, v3 := hash, hash, hash, hash

		for len(ptr) >= 32 {
			v0 += u64s(ptr[:8]) * metroK0
			v0 = rol64(v0, -29) + v2
			v1 += u64s(ptr[8:16]) * metroK1
			v1 = rol64(v1, -29) + v3
			v2 += u64s(ptr[16:24]) * metroK2
			v2 = rol64(v2, -29) + v0
			v3 += u64s(ptr[24:32]) * metroK3
			v3 = rol64(v3, -29) + v1
			ptr = ptr[32:]
		}

		v2 ^= rol64(((v0+v3)*metroK0)+v1, -37) * metroK1
		v3 ^= rol64(((v1+v2)*metroK1)+v0, -37) * metroK0
		v0 ^= rol64(((v0+v2)*metroK0)+v3, -37) * metroK1
		v1 ^= rol64(((v1+v3)*metroK1)+v2, -37) * metroK0
		hash += v0 ^ v1
	}

	if len(ptr) >= 16 {
		v0 := hash + (u64s(ptr[:8]) * metroK2)
		v0 = rol64(v0, -29) * metroK3
		v1 := hash + (u64s(ptr[8:16]) * metroK2)
		v1 = rol64(v1, -29) * metroK3
		v0 ^= rol64(v0*metroK0, -21) + v1
		v1 ^= rol64(v1*metroK3, -21) + v0
		hash += v1
		ptr = ptr[16:]
	}

	if len(ptr) >= 8 {
		hash += u64s(ptr[:8]) * metroK3
		ptr = ptr[8:]
		hash ^= rol64(hash, -55) * metroK1
	}

	if len(ptr) >= 4 {
		hash += uint64(u32s(ptr[:4])) * metroK3
		hash ^= rol64(hash, -26) * metroK1
		ptr = ptr[4:]
	}

	if len(ptr) >= 2 {
		hash += uint64(u16s(ptr[:2])) * metroK3
		ptr = ptr[2:]
		hash ^= rol64(hash, -48) * metroK1
	}

	if len(ptr) >= 1 {
		hash += uint64(ptr[0]) * metroK3
		hash ^= rol64(hash, -37) * metroK1
	}

	hash ^= rol64(hash, -28)
	hash *= metroK0
	hash ^= rol64(hash, -29)

	return hash
}
