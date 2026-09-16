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

package backend

import (
	"fmt"
	"go/ast"

	"github.com/gx-org/backend/dtypes"
	"github.com/gx-org/backend/shape"
)

type (
	// Value in the graph.
	Value interface {
		Graph() Function
	}

	// Tuple bundles multiple Nodes together.
	Tuple interface {
		Value

		// Element returns a Node representing the ith element of the tuple.
		Element(i int) (Value, error)

		// Size returns the number of elements in the tuple.
		Size() int

		// Unpack returns the tuple's constituent Nodes.
		Unpack() ([]Value, error)
	}

	// Executable runs a node in a compiled graph.
	Executable interface {
		Run([]Handle) (out, traces []DeviceHandle, err error)
	}

	// OutputNode is an output node in the graph.
	OutputNode struct {
		Node  Value
		Shape *shape.Shape
	}

	// Function implemented by a backend.
	// The GX interpreter uses this interface to build a graph for the backend.
	Function interface {
		// Platform used by the graph.
		Platform() Platform

		// Graph returns the graph in which the nodes are created into.
		Graph() Function

		// Constant returns a node representing a numerical constant value in the graph.
		Constant(value HostBuffer) (Value, error)

		// NewAtomLiteral creates a node from a constant atom.
		NewAtomLiteral(v any) (Value, error)

		// NewArrayLiteral creates a node from a constant array.
		NewArrayLiteral(flat any, axlengths ...int) (Value, error)

		// Tuple returns a node representing a tuple of nodes.
		Tuple(nodes []Value) (Tuple, error)

		// Call returns a node that invokes a subgraph.
		Call(sg *Subgraph, args ...Value) (Value, error)

		// Subgraph returns a Graph instance that maps to a new subgraph.
		Subgraph(name string, args []*shape.Shape) (Function, error)

		// Argument returns a node set by a caller when calling the function.
		Argument(name string, shape *shape.Shape, index int) (Value, error)

		// LogicalNot returns the logical not of x.
		LogicalNot(x Value) (Value, error)

		// Neg returns the negation of x.
		Neg(x Value) (Value, error)

		// Binary returns a node applying a binary operator between two nodes.
		Binary(op *ast.BinaryExpr, x, y Value) (Value, error)

		// Reshape returns a reshape operator node.
		Reshape(x Value, axisLengths []int) (Value, error)

		// Concat concatenates multiple arrays into a single array.
		Concat(axis int, nodes []Value) (Value, error)

		// Cast returns a cast/convert operator node.
		Cast(x Value, target dtypes.DType) (Value, error)

		// Bitcast casts a byte array into a given data type.
		Bitcast(x Value, target dtypes.DType) (Value, error)

		// Slice returns a slice on a node.
		Slice(x Value, index int) (Value, error)

		// Set returns a node to set a slice in an array.
		Set(x, updates Value, index []Value) (Value, error)

		// Dot product between x and y.
		Dot(x, y Value) (Value, error)

		// DotGeneral returns a general dot operator node.
		DotGeneral(x, y Value, batchAxes, reduceAxes [2][]int) (Value, error)

		// While returns a while loop node.
		While(cond, body *Subgraph, state Value) (Value, error)

		// BroadcastInDim broadcasts data across a given set of axis.
		BroadcastInDim(x Value, shape *shape.Shape, broadcastAxes []int) (Value, error)

		// Iota returns a node filling an array with values from 0 to number of elements-1.
		Iota(sh *shape.Shape, iotaAxis int) (Value, error)

		// ArgMinMax applies a min or max operator over an axis
		ArgMinMax(x Value, axis int, outputDType dtypes.DType, isMin bool) (Value, error)

		// ReduceMax applies max over axes.
		ReduceMax(x Value, axes []int) (Value, error)

		// ReduceSum sums over axes.
		ReduceSum(x Value, axes []int) (Value, error)

		// Split an array along an axis.
		Split(x Value, axis, numSplits int) (Value, error)

		// Gather data from an array.
		Gather(x Value, startIndices Value, indexVectorAxis int, offsetAxes []int, collapsedSliceAxes []int, startIndexMap []int, sliceSizes []int, indicesAreSorted bool) (Value, error)

		// Transpose the array.
		Transpose(x Value, permutation []int) (Value, error)

		// Abs returns the absolute value of x.
		Abs(x Value) (Value, error)
		// Ceil returns the ceiling of x.
		Ceil(x Value) (Value, error)
		// Cos returns the cosine of x.
		Cos(x Value) (Value, error)
		// Erf returns the error function of x.
		Erf(x Value) (Value, error)
		// Exp returns the exponential of x.
		Exp(x Value) (Value, error)
		// Expm1 returns Exp(x)-1.
		Expm1(x Value) (Value, error)
		// Floor returns the floor of x.
		Floor(x Value) (Value, error)
		// Log returns the natural logarithm of x.
		Log(x Value) (Value, error)
		// Log1p returns log(1+x).
		Log1p(x Value) (Value, error)
		// Logistic returns 1/(1+exp(-x)).
		Logistic(x Value) (Value, error)
		// Min returns the minimum between x and y.
		Min(x, y Value) (Value, error)
		// Max returns the maximum between x and y.
		Max(x, y Value) (Value, error)
		// Pow returns x to the power of y.
		Pow(x, y Value) (Value, error)
		// Round returns the nearest integer of x.
		Round(x Value) (Value, error)
		// Rsqrt returns 1/sqrt(x).
		Rsqrt(x Value) (Value, error)
		// Sign returns the sign of x.
		Sign(x Value) (Value, error)
		// Sin returns the sine of x.
		Sin(x Value) (Value, error)
		// Sqrt returns sqrt(x).
		Sqrt(x Value) (Value, error)
		// Tanh returns the hyperbolic tangent of x.
		Tanh(x Value) (Value, error)

		// RngBitGenerator generates random values of the given shape using the provided RNG state.
		// It returns the updated RNG state and the generated values.
		RngBitGenerator(state Value, shape *shape.Shape) (Value, Value, error)
	}

	// Subgraph bundles a Graph and its output node together.
	Subgraph struct {
		Graph  Function
		Result OutputNode
	}
)

// String representation of an output node.
func (out *OutputNode) String() string {
	return fmt.Sprintf("%s: %v", out.Shape.String(), out.Node)
}
