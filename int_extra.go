// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implement extra functionality not in big.Int.
//
// Because the underlying C compiler may use 32-bit longs we
// use 32-bit integers here for maximum portability and speed.

package gmp

/*
#cgo LDFLAGS: -lgmp
#include <gmp.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// Sqrt sets x to the truncated integer part of the square root of x
//
// NB This is not part of big.Int
func (z *Int) Sqrt(x *Int) *Int {
	x.doinit()
	z.doinit()
	C.mpz_sqrt((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)))
	return z
}

// Swap exchanges the values of z and x and returns the new z.
//
// NB This is not part of big.Int
func (z *Int) Swap(x *Int) *Int {
	x.doinit()
	z.doinit()
	C.mpz_swap((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)))
	return z
}

// AddUint32 sets z to the sum x+y and returns z.
//
// NB This is not part of big.Int
func (z *Int) AddUint32(x *Int, y uint32) *Int {
	x.doinit()
	z.doinit()
	C.mpz_add_ui((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), C.ulong(y))
	return z
}

// SubUint32 sets z to the difference x-y and returns z.
//
// NB This is not part of big.Int
func (z *Int) SubUint32(x *Int, y uint32) *Int {
	x.doinit()
	z.doinit()
	C.mpz_sub_ui((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), C.ulong(y))
	return z
}

// Uint32Sub sets z to the difference x-y and returns z.
//
// NB This is not part of big.Int
func (z *Int) Uint32Sub(x uint32, y *Int) *Int {
	y.doinit()
	z.doinit()
	C.mpz_ui_sub((C.mpz_ptr)(unsafe.Pointer(&z.i)), C.ulong(x), (C.mpz_ptr)(unsafe.Pointer(&y.i)))
	return z
}

// MulUint32 sets z to the product x*y and returns z.
//
// NB This is not part of big.Int
func (z *Int) MulUint32(x *Int, y uint32) *Int {
	x.doinit()
	z.doinit()
	C.mpz_mul_ui((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), C.ulong(y))
	return z
}

// MulInt32 sets z to the product x*y and returns z.
//
// NB This is not part of big.Int
func (z *Int) MulInt32(x *Int, y int32) *Int {
	x.doinit()
	z.doinit()
	C.mpz_mul_si((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), C.long(y))
	return z
}

// AddMul sets z to z + x*y and returns z.
//
// NB This is not part of big.Int
func (z *Int) AddMul(x, y *Int) *Int {
	x.doinit()
	y.doinit()
	z.doinit()
	C.mpz_addmul((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), (C.mpz_ptr)(unsafe.Pointer(&y.i)))
	return z
}

// AddMulUint32 sets z to z + x*y and returns z.
//
// NB This is not part of big.Int
func (z *Int) AddMulUint32(x *Int, y uint32) *Int {
	x.doinit()
	z.doinit()
	C.mpz_addmul_ui((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), C.ulong(y))
	return z
}

// SubMul sets z to z - x*y and returns z.
//
// NB This is not part of big.Int
func (z *Int) SubMul(x, y *Int) *Int {
	x.doinit()
	y.doinit()
	z.doinit()
	C.mpz_submul((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), (C.mpz_ptr)(unsafe.Pointer(&y.i)))
	return z
}

// SubMulUint32 sets z to z - x*y and returns z.
//
// NB This is not part of big.Int
func (z *Int) SubMulUint32(x *Int, y uint32) *Int {
	x.doinit()
	z.doinit()
	C.mpz_submul_ui((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i)), C.ulong(y))
	return z
}

// compared reduces the result of a GMP comparison to one of {-1, 0, 1}.
func compared(i C.int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	} else {
		return 0
	}
}

// CmpUint32 compares z and x and returns:
//
//   -1 if z <  x
//    0 if z == x
//   +1 if z >  x
//
// NB This is not part of big.Int
func (z *Int) CmpUint32(x uint32) int {
	z.doinit()
	return compared(C._mpz_cmp_ui((C.mpz_ptr)(unsafe.Pointer(&z.i)), C.ulong(x)))
}

// CmpInt32 compares z and x and returns:
//
//   -1 if z <  x
//    0 if z == x
//   +1 if z >  x
//
// NB This is not part of big.Int
func (z *Int) CmpInt32(x int32) int {
	z.doinit()
	return compared(C._mpz_cmp_si((C.mpz_ptr)(unsafe.Pointer(&z.i)), C.long(x)))
}

// CmpAbs compares |z| and |x| and returns:
//
//   -1 if |z| <  |x|
//    0 if |z| == |x|
//   +1 if |z| >  |x|
//
// NB This is not part of big.Int
func (z *Int) CmpAbs(x *Int) int {
	x.doinit()
	z.doinit()
	return compared(C.mpz_cmpabs((C.mpz_ptr)(unsafe.Pointer(&z.i)), (C.mpz_ptr)(unsafe.Pointer(&x.i))))
}

// CmpAbsUint32 compares |z| and |x| and returns:
//
//   -1 if |z| <  |x|
//    0 if |z| == |x|
//   +1 if |z| >  |x|
//
// NB This is not part of big.Int
func (z *Int) CmpAbsUint32(x uint32) int {
	z.doinit()
	return compared(C.mpz_cmpabs_ui((C.mpz_ptr)(unsafe.Pointer(&z.i)), C.ulong(x)))
}

// Uint32 returns the uint32 representation of z, if z fits into a uint32.
// If z is too big then the least significant bits that do fit are returned.
// The sign of z is ignored, only the absolute value is used.
//
// NB This is not part of big.Int
func (z *Int) Uint32() uint32 {
	z.doinit()
	return uint32(C.mpz_get_ui((C.mpz_ptr)(unsafe.Pointer(&z.i))))
}

// Int32 returns the int32 representation of z, if z fits into a signed int32.
// If z is too big to fit in a int32 then the result is undefined.
//
// NB This is not part of big.Int
func (z *Int) Int32() int32 {
	z.doinit()
	return int32(C.mpz_get_si((C.mpz_ptr)(unsafe.Pointer(&z.i))))
}
