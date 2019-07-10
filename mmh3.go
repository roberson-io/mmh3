package mmh3

import (
	"encoding/binary"
	"github.com/roberson-io/mmh3/MurmurHash3"
	"unsafe"
)

// Endianness checking pattern from TensorFlow
// https://github.com/tensorflow/tensorflow/blob/master/tensorflow/go/tensor.go#L488-L505
var nativeEndian binary.ByteOrder

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		panic("Could not determine native endianness.")
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
            block := nativeEndian.Uint32(key[i : i+blockSize])
            data32 = append(data32, block)
        }
        ptrKey = uintptr(unsafe.Pointer(&data32[0]))
    } else {
        blockSize = 8
        nBytes = 16
        for i := 0; i < len(key); i += blockSize {
            block := nativeEndian.Uint64(key[i : i+blockSize])
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

func Hash_x86_32(key []byte, seed uint32) []byte {
    is32Bit := true
    return mmh3Hash(key, seed, is32Bit, MurmurHash3.MurmurHash3_x86_32)
}

func Hash_x86_128(key []byte, seed uint32) []byte {
    is32Bit := false
    return mmh3Hash(key, seed, is32Bit, MurmurHash3.MurmurHash3_x86_128)
}

func Hash_x64_128(key []byte, seed uint32) []byte {
    is32Bit := false
    return mmh3Hash(key, seed, is32Bit, MurmurHash3.MurmurHash3_x64_128)
}
