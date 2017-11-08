// Package hashx implements allocation-free standard hashing functions
package hashx

func u64s(s string, offset int) uint64 {
	return uint64(s[offset+0]) | uint64(s[offset+1])<<8 | uint64(s[offset+2])<<16 | uint64(s[offset+3])<<24 | uint64(s[offset+4])<<32 | uint64(s[offset+5])<<40 | uint64(s[offset+6])<<48 | uint64(s[offset+7])<<56
}

func u32s(s string, offset int) uint32 {
	return uint32(s[offset+0]) | uint32(s[offset+1])<<8 | uint32(s[offset+2])<<16 | uint32(s[offset+3])<<24
}

func u64(buf []byte) uint64 {
	// go compiler recognizes this pattern and optimizes it on little endian platforms
	return uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 | uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56
}

func u32(buf []byte) uint32 {
	return uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24
}

func fmix(k uint64) uint64 {
	k ^= k >> 33
	k *= 0xff51afd7ed558ccd
	k ^= k >> 33
	k *= 0xc4ceb9fe1a85ec53
	k ^= k >> 33
	return k
}

func rol64_1(u uint64) uint64 {
	return u<<1 | u>>63
}
func rol64_7(u uint64) uint64 {
	return u<<7 | u>>57
}
func rol64_11(u uint64) uint64 {
	return u<<11 | u>>53
}
func rol64_12(u uint64) uint64 {
	return u<<12 | u>>52
}
func rol64_18(u uint64) uint64 {
	return u<<18 | u>>46
}
func rol64_23(u uint64) uint64 {
	return u<<23 | u>>41
}
func rol64_27(u uint64) uint64 {
	return u<<27 | u>>37
}
func rol64_31(u uint64) uint64 {
	return u<<31 | u>>33
}
func rol64_33(u uint64) uint64 {
	return u<<33 | u>>31
}

func rol32_1(u uint32) uint32 {
	return u<<1 | u>>31
}
func rol32_7(u uint32) uint32 {
	return u<<7 | u>>25
}
func rol32_11(u uint32) uint32 {
	return u<<11 | u>>21
}
func rol32_12(u uint32) uint32 {
	return u<<12 | u>>20
}
func rol32_13(u uint32) uint32 {
	return u<<13 | u>>19
}
func rol32_15(u uint32) uint32 {
	return u<<15 | u>>17
}
func rol32_17(u uint32) uint32 {
	return u<<17 | u>>15
}
func rol32_18(u uint32) uint32 {
	return u<<18 | u>>14
}
