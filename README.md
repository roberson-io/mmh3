# MurmurHash3 for Golang
This is a Golang wrapper for the original [MurmurHash3](https://github.com/aappleby/smhasher/blob/master/src/MurmurHash3.cpp) C++ implementation. It was created with [SWIG](http://swig.org). The goal is to make something similar to the Python [mmh3](https://github.com/hajimes/mmh3) package for Golang.

## Usage
You probably just want to use the three helper functions in the mmh3.go file at the root of this package:
- `Hashx86_32`
- `Hashx86_128`
- `Hashx64_128`

The files in the `MurmurHash3` directory are for SWIG and are not very readable.

The Python [mmh3](https://github.com/hajimes/mmh3) package has a `hash` function that is just shorthand for the x86 32-bit MurmurHash3 function.  It also returns signed integers by default. I have done no such thing in this package. You must pass the input as a byte slice, provide a seed value, and specify which MurmurHash3 function you want to use. If you want to hash a string a la `print(hex(mmh3.hash('foo')))` in Python mmh3, you might do something like:

```golang
import (
    "encoding/binary"
    "fmt"
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
I am aware of DataDog's Golang implementation of [mmh3](https://github.com/DataDog/mmh3). Key differences between that project and this one are:
- DataDog's implementation is pure Go while this one is just a wrapper around the original C++ implementation.
- The DataDog implementation does not allow you to specify a seed value.