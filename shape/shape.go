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

// Package shape defines the shape of an array.
package shape

import (
	"fmt"
	"strings"

	"github.com/gx-org/backend/dtype"
)

// Shape represents the shape of an array, that is the datatype of the
// elements stored in the array and a list of axis lengths in major-to-minor
// order.
type Shape struct {
	DType       dtype.DataType
	AxisLengths []int
}

// OuterAxisLength returns the shape's outermost axis length, or 1 for rank-0 shapes.
func (s Shape) OuterAxisLength() int {
	if len(s.AxisLengths) == 0 {
		return 1
	}
	return s.AxisLengths[0]
}

// IsAtomic returns true for the shape of an atomic value, that is
// a single value with no axis.
func (s Shape) IsAtomic() bool {
	return len(s.AxisLengths) == 0
}

// Size returns the number of elements of DType are needed for this shape. It's the product of all dimensions.
func (s Shape) Size() int {
	size := 1
	for _, d := range s.AxisLengths {
		size *= d
	}
	return size
}

// ByteSize returns the size of the buffer, in bytes, to store the data specified by the shape.
func (s Shape) ByteSize() int {
	return dtype.Sizeof(s.DType) * s.Size()
}

func (s Shape) String() string {
	axes := make([]string, len(s.AxisLengths))
	for i, axisLength := range s.AxisLengths {
		axes[i] = fmt.Sprintf("[%d]", axisLength)
	}
	return strings.Join(axes, "") + s.DType.String()
}

// ArrayI is a minimum generic array interface.
type ArrayI[T dtype.GoDataType] interface {
	// Shape returns the size of all the axes of the array.
	Shape() []int

	// Flat returns the data stored by the array.
	// The length of the returned slice should match the size of the shape.
	Flat() []T
}
