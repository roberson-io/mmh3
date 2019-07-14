package mmh3

import (
	"encoding/binary"
	"unsafe"

	"github.com/roberson-io/mmh3/MurmurHash3"
)

var endian binary.ByteOrder

func init() {
	twoBytes := [2]byte{}
	*(*uint16)(unsafe.Pointer(&twoBytes[0])) = uint16(0x1234)

	switch twoBytes {
	case [2]byte{0x12, 0x34}:
		endian = binary.BigEndian
	case [2]byte{0x34, 0x12}:
		endian = binary.LittleEndian
	default:
		panic("Unknown endianness.")
	}
}

type mmh3Function func(uintptr, int, MurmurHash3.Uint32_t, uintptr)

func mmh3Hash(key []byte, seed uint32, is32Bit bool, mmh3Func mmh3Function) []byte {
	var data32 []uint32
	var data64 []uint64
	var blockSize, nBytes int
	var ptrKey uintptr
	if is32Bit {
		blockSize = 4
		nBytes = 4
		for i := 0; i < len(key); i += blockSize {
			block := endian.Uint32(key[i : i+blockSize])
			data32 = append(data32, block)
		}
		ptrKey = uintptr(unsafe.Pointer(&data32[0]))
	} else {
		blockSize = 8
		nBytes = 16
		for i := 0; i < len(key); i += blockSize {
			block := endian.Uint64(key[i : i+blockSize])
			data64 = append(data64, block)
		}
		ptrKey = uintptr(unsafe.Pointer(&data64[0]))
	}
	ptrSeed := MurmurHash3.SwigcptrUint32_t(unsafe.Pointer(&seed))
	hash := make([]byte, nBytes)
	ptrHash := uintptr(unsafe.Pointer(&hash[0]))
	mmh3Func(ptrKey, len(key), ptrSeed, ptrHash)
	return hash
}

// Hashx86_32 is a helper function for MurmurHash3_x86_32 in the original C++
// implementation. It accepts data to be hashed as a byte slice as well as a
// seed value. It returns a byte slice with the hash.
func Hashx86_32(key []byte, seed uint32) []byte {
	is32Bit := true
	return mmh3Hash(key, seed, is32Bit, MurmurHash3.MurmurHash3_x86_32)
}

// Hashx86_128 is a helper function for MurmurHash3_x86_128 in the original C++
// implementation. It accepts data to be hashed as a byte slice as well as a
// seed value. It returns a byte slice with the hash.
func Hashx86_128(key []byte, seed uint32) []byte {
	is32Bit := false
	return mmh3Hash(key, seed, is32Bit, MurmurHash3.MurmurHash3_x86_128)
}

// Hashx64_128 is a helper function for MurmurHash3_x64_128 in the original C++
// implementation. It accepts data to be hashed as a byte slice as well as a
// seed value. It returns a byte slice with the hash.
func Hashx64_128(key []byte, seed uint32) []byte {
	is32Bit := false
	return mmh3Hash(key, seed, is32Bit, MurmurHash3.MurmurHash3_x64_128)
}
