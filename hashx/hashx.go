// Package hashx implements allocation-free standard hashing functions
package hashx

import (
	"encoding/binary"
	"math/bits"
)

func u64s(s string) uint64 {
	return uint64(s[0]) | uint64(s[1])<<8 | uint64(s[2])<<16 | uint64(s[3])<<24 | uint64(s[4])<<32 | uint64(s[5])<<40 | uint64(s[6])<<48 | uint64(s[7])<<56
}

func u32s(s string) uint32 {
	return uint32(s[0]) | uint32(s[1])<<8 | uint32(s[2])<<16 | uint32(s[3])<<24
}

func u16s(s string) uint32 {
	return uint32(s[0]) | uint32(s[1])<<8
}

func u64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

func u32(buf []byte) uint32 {
	return binary.LittleEndian.Uint32(buf)
}

func u16(buf []byte) uint16 {
	return binary.LittleEndian.Uint16(buf)
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
