// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package dtype defines data types that can be supported by a platform.
package dtype

import (
	"fmt"
	"reflect"
	"unsafe"
)

// DataType is the type of an atomic value or type of the data stored in an array.
type DataType uint32

// All supported types
const (
	Invalid DataType = iota

	Bool
	Int
	Int32
	Int64
	Uint32
	Uint64
	Float32
	Float64

	MaxDataType = 1 << 16 // Maximum value for a datatype.
)

// String returns a string representation of a kind.
func (dt DataType) String() string {
	switch dt {
	case Bool:
		return "bool"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case Uint32:
		return "uint32"
	case Uint64:
		return "uint64"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	}
	return "invalid"
}

// Float is a constraint supporting floating-point type.
type Float interface {
	~float32 | ~float64
}

// Signed is a constraint supporting signed integer type.
type Signed interface {
	~int32 | ~int64
}

// Unsigned is a constraint supporting unsigned integer type.
type Unsigned interface {
	~uint32 | ~uint64
}

// NonAlgebraType are types on which common algebra operations are NOT supported.
type NonAlgebraType interface {
	~bool
}

// IntegerType are types on which integer algebra operations are supported.
type IntegerType interface {
	Signed | Unsigned
}

// AlgebraType are types on which common algebra operations between integers and floats are supported.
type AlgebraType interface {
	Float | IntegerType
}

// GoDataType that can be stored in an array.
type GoDataType interface {
	AlgebraType | NonAlgebraType
}

// Generic returns a dtype from a Go type.
func Generic[T GoDataType]() DataType {
	var t T
	switch (any(t)).(type) {
	case bool:
		return Bool
	case float32:
		return Float32
	case float64:
		return Float64
	case int32:
		return Int32
	case int64:
		return Int64
	case uint32:
		return Uint32
	case uint64:
		return Uint64
	}
	return Invalid
}

// Sizes of data type (in bytes).
const (
	BoolSize    = 1
	Int32Size   = 4
	Int64Size   = 8
	Uint32Size  = 4
	Uint64Size  = 8
	Float32Size = 4
	Float64Size = 8
)

// Sizeof returns the size of an atomic value of a data type.
func Sizeof(dt DataType) int {
	switch dt {
	case Bool:
		return BoolSize
	case Int32:
		return Int32Size
	case Int64:
		return Int64Size
	case Uint32:
		return Uint32Size
	case Uint64:
		return Uint64Size
	case Float32:
		return Float32Size
	case Float64:
		return Float64Size
	}
	panic("invalid datatype")
}

// ToSlice converts a []byte buffer into a slice of a given Go type.
func ToSlice[T GoDataType](data []byte) []T {
	var t T
	size := int(unsafe.Sizeof(t))
	if len(data)%size != 0 {
		typeName := reflect.TypeFor[T]().String()
		panic(fmt.Sprintf("data [%d]byte cannot be casted to []%s: %d %% sizeof(%s) != 0", len(data), typeName, len(data), typeName))
	}
	length := len(data) / size
	return unsafe.Slice((*T)(unsafe.Pointer(&data[0])), length)
}
