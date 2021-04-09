// Package hashx implements allocation-free standard hashing functions
package hashx

import "math/bits"

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

func rol64(u uint64, k int) uint64 { return bits.RotateLeft64(u, k) }
func rol32(u uint32, k int) uint32 { return bits.RotateLeft32(u, k) }
