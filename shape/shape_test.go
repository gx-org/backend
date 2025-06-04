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
	"testing"

	"github.com/gx-org/backend/dtype"
)

func TestShapeEqual(t *testing.T) {
	tests := []struct {
		desc         string
		x, is, isNot Shape
	}{
		{
			desc:  "dtype",
			x:     Shape{DType: dtype.Float32},
			is:    Shape{DType: dtype.Float32},
			isNot: Shape{DType: dtype.Int32},
		},
		{
			desc:  "number of axis",
			x:     Shape{DType: dtype.Float32, AxisLengths: []int{1, 2}},
			is:    Shape{DType: dtype.Float32, AxisLengths: []int{1, 2}},
			isNot: Shape{DType: dtype.Float32, AxisLengths: []int{1}},
		},
		{
			desc:  "axis lengths",
			x:     Shape{DType: dtype.Float32, AxisLengths: []int{1, 2}},
			is:    Shape{DType: dtype.Float32, AxisLengths: []int{1, 2}},
			isNot: Shape{DType: dtype.Float32, AxisLengths: []int{1, 3}},
		},
	}
	for i, test := range tests {
		if !test.x.Equal(&test.is) {
			t.Errorf("test %d:%s: %v == %v is false", i, test.desc, test.x, test.is)
		}
		if test.x.Equal(&test.isNot) {
			t.Errorf("test %d:%s: %v != %v is false", i, test.desc, test.x, test.isNot)
		}
	}
}
