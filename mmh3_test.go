package mmh3

import (
	"testing"
)

func TestHashx86_32(t *testing.T) {
	known := map[string]map[uint32]uint32{
		"a": {
			0:          0x3c2569b2,
			1:          0x588adce8,
			0xffffffff: 0x2a684527,
		},
		"Hello, world!": {
			0:          0xc0363e43,
			1:          0xaa5dc85b,
			0xffffffff: 0x07d2b7b4,
		},
	}
	for message, seeds := range known {
		for seed, expected := range seeds {
			key := []byte(message)
			hash := Hashx86_32(key, seed)
			if endian.Uint32(hash) != expected {
				t.Errorf(
					"Hashx86_32: key: %s seed: %x hash: %x expected: %x",
					message,
					seed,
					endian.Uint32(hash),
					expected,
				)
			}
		}
	}
}

func TestHashx86_128(t *testing.T) {
	known := map[string]map[uint32][2]uint64{
		"a": {
			0:          {0x5556b01ba794933c, 0x5556b01b5556b01b},
			1:          {0x1c5281c549db9b04, 0x1c5281c51c5281c5},
			0xffffffff: {0x90f0af0878bc24b1, 0x90f0af0890f0af08},
		},
		"Hello, world!": {
			0:          {0xf0638dfc26acdba7, 0xafdd4c3402b4263},
			1:          {0xd8b061d317e7f784, 0x92b0b92caa180c8b},
			0xffffffff: {0xcbbc52238d191775, 0x461b3d6b30c75e10},
		},
	}
	for message, seeds := range known {
		for seed, expected := range seeds {
			key := []byte(message)
			hash := Hashx86_128(key, seed)
			convertedHash := [2]uint64{endian.Uint64(hash[:8]), endian.Uint64(hash[8:])}
			if convertedHash != expected {
				t.Errorf(
					"Hashx86_128: key: %s seed: %x hash: [%x, %x] expected: %x",
					message,
					seed,
					endian.Uint64(hash[:8]),
					endian.Uint64(hash[8:]),
					expected,
				)
			}
		}
	}
}

func TestHashx64_128(t *testing.T) {
	known := map[string]map[uint32][2]uint64{
		"a": {
			0:          {0x85555565f6597889, 0xe6b53a48510e895a},
			1:          {0x47eae1073748cf70, 0x6be0518ad2ed3728},
			0xffffffff: {0xbef385faead16340, 0xa9363d237b2ee74c},
		},
		"Hello, world!": {
			0:          {0xf1512dd1d2d665df, 0x2c326650a8f3c564},
			1:          {0xf2f417beaba4a824, 0x2f783a2a5a9049ee},
			0xffffffff: {0x26d7b85c4f149498, 0x318ddf746ca6b8c7},
		},
	}
	for message, seeds := range known {
		for seed, expected := range seeds {
			key := []byte(message)
			hash := Hashx64_128(key, seed)
			convertedHash := [2]uint64{endian.Uint64(hash[:8]), endian.Uint64(hash[8:])}
			if convertedHash != expected {
				t.Errorf(
					"Hashx64_128: key: %s seed: %x hash: [%x, %x] expected: %x",
					message,
					seed,
					endian.Uint64(hash[:8]),
					endian.Uint64(hash[8:]),
					expected,
				)
			}
		}
	}
}
