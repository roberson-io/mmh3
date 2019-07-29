[![BuildStatus](https://travis-ci.org/roberson-io/mmh3.svg?branch=master)](https://travis-ci.org/roberson-io/mmh3) [![GoDoc](https://godoc.org/github.com/roberson-io/mmh3?status.svg)](http://godoc.org/github.com/roberson-io/mmh3)


# MurmurHash3 for Golang
This is a Golang wrapper for the original [MurmurHash3](https://github.com/aappleby/smhasher/blob/master/src/MurmurHash3.cpp) C++ implementation. It was created with [SWIG](http://swig.org). The goal is to make something similar to the Python [mmh3](https://github.com/hajimes/mmh3) package for Golang.

## Usage
You probably just want to use the three helper functions in the mmh3.go file at the root of this package:
- `Hashx86_32`
- `Hashx86_128`
- `Hashx64_128`

The files in the `MurmurHash3` directory are for SWIG and are not very readable.

The Python [mmh3](https://github.com/hajimes/mmh3) package has a `hash` function that is just shorthand for the x86 32-bit MurmurHash3 function.  It also returns signed integers by default. I have done no such thing in this package. You must pass the input as a byte slice, provide a seed value, and specify which MurmurHash3 function you want to use. If you want to hash a string à la `print(hex(mmh3.hash('foo')))` in Python mmh3, you might do something like:

```golang
import (
    "encoding/binary"
    "fmt"

    "github.com/roberson-io/mmh3"
)

 key := []byte("foo")
 var seed uint32 = 0
 hash := mmh3.Hashx86_32(key, seed)
 fmt.Printf("%x\n", binary.LittleEndian.Uint32(hash))
```

You can call the other two functions the same way:
```golang
 key := []byte("foo")
 var seed uint32 = 0

 hashx86 := mmh3.Hashx86_128(key, seed)
 fmt.Printf(
     "%x\t%x\n",
     binary.LittleEndian.Uint64(hashx86[:8]),
     binary.LittleEndian.Uint64(hashx86[8:]),
 )

 hashx64 := mmh3.Hashx64_128(key, seed)
 fmt.Printf(
     "%x\t%x\n",
     binary.LittleEndian.Uint64(hashx64[:8]),
     binary.LittleEndian.Uint64(hashx64[8:]),
 )
```

## Why?
I am aware of reusee's Golang implementation of [mmh3](https://github.com/reusee/mmh3). Key differences between that project and this one are:
- This just a wrapper around the original C++ implementation (not pure Go)
- This allows you to specify a seed value.
- This does not implement the hash.Hash interface used in standard library packages (crypto/md5, crypto/sha256, etc.).
