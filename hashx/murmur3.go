package hashx

import "encoding/binary"

const (
	mm32C1 uint32 = 0xcc9e2d51
	mm32C2 uint32 = 0x1b873593

	mm128C1 = 0x87c37b91114253d5
	mm128C2 = 0x4cf5ad432745937f
)

// MM32 calculates 32-bit murmur3 hash.
func MM32(data []byte) (h1 uint32) {
	nblocks := len(data) / 4
	for b := 0; b < nblocks; b++ {
		k1 := binary.LittleEndian.Uint32(data[b*4:])

		k1 *= mm32C1
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= mm32C2

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19) // rotl32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	tail := data[nblocks*4:]

	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= mm32C1
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= mm32C2
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}

// MM32String calculates 32-bit murmur3 hash.
func MM32String(data string) (h1 uint32) {
	nblocks := len(data) / 4
	for b := 0; b < nblocks; b++ {
		k1 := atou32(data, b*4)

		k1 *= mm32C1
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= mm32C2

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19) // rotl32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	hlen := nblocks * 4

	var k1 uint32
	switch (len(data) - hlen) & 3 {
	case 3:
		k1 ^= uint32(data[hlen+2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(data[hlen+1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(data[hlen+0])
		k1 *= mm32C1
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= mm32C2
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}

// MM64 calculates 64-bit murmur3 hash.
func MM64(data []byte) uint64 {
	h1, _ := MM128(data)
	return h1
}

// MM64String calculates 64-bit murmur3 hash.
func MM64String(data string) uint64 {
	h1, _ := MM128String(data)
	return h1
}

// MM128 calculates 128-bit murmur3 hash.
func MM128(data []byte) (h1, h2 uint64) {
	nblocks := len(data) / 16
	for b := 0; b < nblocks; b++ {
		k1 := binary.LittleEndian.Uint64(data[b*16:])
		k2 := binary.LittleEndian.Uint64(data[b*16+8:])

		k1 *= mm128C1
		k1 = (k1 << 31) | (k1 >> 33) // rotl64(k1, 31)
		k1 *= mm128C2
		h1 ^= k1

		h1 = (h1 << 27) | (h1 >> 37) // rotl64(h1, 27)
		h1 += h2
		h1 = h1*5 + 0x52dce729

		k2 *= mm128C2
		k2 = (k2 << 33) | (k2 >> 31) // rotl64(k2, 33)
		k2 *= mm128C1
		h2 ^= k2

		h2 = (h2 << 31) | (h2 >> 33) // rotl64(h2, 31)
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}

	tail := data[nblocks*16:]

	var k1, k2 uint64
	switch len(tail) & 15 {
	case 15:
		k2 ^= uint64(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(tail[8]) << 0

		k2 *= mm128C2
		k2 = (k2 << 33) | (k2 >> 31) // rotl64(k2, 33)
		k2 *= mm128C1
		h2 ^= k2

		fallthrough
	case 8:
		k1 ^= uint64(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(tail[0]) << 0
		k1 *= mm128C1
		k1 = (k1 << 31) | (k1 >> 33) // rotl64(k1, 31)
		k1 *= mm128C2
		h1 ^= k1
	}

	h1 ^= uint64(len(data))
	h2 ^= uint64(len(data))

	h1 += h2
	h2 += h1

	h1 = fmix64(h1)
	h2 = fmix64(h2)

	h1 += h2
	h2 += h1

	return
}

// MM128String calculates 128-bit murmur3 hash.
func MM128String(data string) (h1, h2 uint64) {
	nblocks := len(data) / 16
	for b := 0; b < nblocks; b++ {
		k1 := atou64(data, b*16)
		k2 := atou64(data, b*16+8)

		k1 *= mm128C1
		k1 = (k1 << 31) | (k1 >> 33) // rotl64(k1, 31)
		k1 *= mm128C2
		h1 ^= k1

		h1 = (h1 << 27) | (h1 >> 37) // rotl64(h1, 27)
		h1 += h2
		h1 = h1*5 + 0x52dce729

		k2 *= mm128C2
		k2 = (k2 << 33) | (k2 >> 31) // rotl64(k2, 33)
		k2 *= mm128C1
		h2 ^= k2

		h2 = (h2 << 31) | (h2 >> 33) // rotl64(h2, 31)
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}

	hlen := nblocks * 16

	var k1, k2 uint64
	switch (len(data) - hlen) & 15 {
	case 15:
		k2 ^= uint64(data[hlen+14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(data[hlen+13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(data[hlen+12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(data[hlen+11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(data[hlen+10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(data[hlen+9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(data[hlen+8]) << 0

		k2 *= mm128C2
		k2 = (k2 << 33) | (k2 >> 31) // rotl64(k2, 33)
		k2 *= mm128C1
		h2 ^= k2

		fallthrough
	case 8:
		k1 ^= uint64(data[hlen+7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(data[hlen+6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(data[hlen+5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(data[hlen+4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(data[hlen+3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(data[hlen+2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(data[hlen+1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(data[hlen+0]) << 0
		k1 *= mm128C1
		k1 = (k1 << 31) | (k1 >> 33) // rotl64(k1, 31)
		k1 *= mm128C2
		h1 ^= k1
	}

	h1 ^= uint64(len(data))
	h2 ^= uint64(len(data))

	h1 += h2
	h2 += h1

	h1 = fmix64(h1)
	h2 = fmix64(h2)

	h1 += h2
	h2 += h1

	return
}

func fmix64(k uint64) uint64 {
	k ^= k >> 33
	k *= 0xff51afd7ed558ccd
	k ^= k >> 33
	k *= 0xc4ceb9fe1a85ec53
	k ^= k >> 33
	return k
}

func atou64(s string, offset int) uint64 {
	return uint64(s[offset+0]) | uint64(s[offset+1])<<8 | uint64(s[offset+2])<<16 | uint64(s[offset+3])<<24 |
		uint64(s[offset+4])<<32 | uint64(s[offset+5])<<40 | uint64(s[offset+6])<<48 | uint64(s[offset+7])<<56
}

func atou32(s string, offset int) uint32 {
	return uint32(s[offset+0]) | uint32(s[offset+1])<<8 | uint32(s[offset+2])<<16 | uint32(s[offset+3])<<24
}
