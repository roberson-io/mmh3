/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: MurmurHash3.i

package MurmurHash3

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_MurmurHash3_9d0bcf074dc0c997(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_MurmurHash3_9d0bcf074dc0c997(swig_intgo arg1);
extern void _wrap_MurmurHash3_x86_32_MurmurHash3_9d0bcf074dc0c997(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3, uintptr_t arg4);
extern void _wrap_MurmurHash3_x86_128_MurmurHash3_9d0bcf074dc0c997(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3, uintptr_t arg4);
extern void _wrap_MurmurHash3_x64_128_MurmurHash3_9d0bcf074dc0c997(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3, uintptr_t arg4);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_MurmurHash3_9d0bcf074dc0c997(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_MurmurHash3_9d0bcf074dc0c997(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func MurmurHash3_x86_32(arg1 uintptr, arg2 int, arg3 Uint32_t, arg4 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	_swig_i_3 := arg4
	C._wrap_MurmurHash3_x86_32_MurmurHash3_9d0bcf074dc0c997(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))
}

func MurmurHash3_x86_128(arg1 uintptr, arg2 int, arg3 Uint32_t, arg4 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	_swig_i_3 := arg4
	C._wrap_MurmurHash3_x86_128_MurmurHash3_9d0bcf074dc0c997(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))
}

func MurmurHash3_x64_128(arg1 uintptr, arg2 int, arg3 Uint32_t, arg4 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	_swig_i_3 := arg4
	C._wrap_MurmurHash3_x64_128_MurmurHash3_9d0bcf074dc0c997(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))
}


type SwigcptrUint32_t uintptr
type Uint32_t interface {
	Swigcptr() uintptr;
}
func (p SwigcptrUint32_t) Swigcptr() uintptr {
	return uintptr(p)
}

