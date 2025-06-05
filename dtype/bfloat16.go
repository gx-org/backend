package dtype

import (
	"math"
	"strconv"
)

// Bfloat16T is a shortened (16-bit) version of the 32-bit IEEE 754 single-precision floating-point
// format (binary32). This implementation only supports conversion to/from other formats, but no
// arithmetic.
//
// Based on the implementation in https://github.com/gomlx/gopjrt, which is in turn derived from
// https://github.com/x448/float16.
type Bfloat16T uint16

// BFloat16FromFloat32 converts a float32 to a BFloat16.
func BFloat16FromFloat32(x float32) Bfloat16T {
	return Bfloat16T(math.Float32bits(x) >> 16)
}

// BFloat16FromFloat64 converts a float32 to a BFloat16.
func BFloat16FromFloat64(x float64) Bfloat16T {
	return BFloat16FromFloat32(float32(x))
}

// Float32 returns a BFloat16 value in float32 format.
func (f Bfloat16T) Float32() float32 {
	return math.Float32frombits(uint32(f) << 16)
}

// Bits convert BFloat16 to an uint16.
func (f Bfloat16T) Bits() uint16 {
	return uint16(f)
}

// String implements fmt.Stringer, and prints a float representation of a BFloat16.
func (f Bfloat16T) String() string {
	return strconv.FormatFloat(float64(f.Float32()), 'f', -1, 32)
}
