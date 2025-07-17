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

// Package ops defines the operations backend can build.
package ops

import (
	"fmt"
	"go/ast"

	"github.com/gx-org/backend/dtype"
	"github.com/gx-org/backend/platform"
	"github.com/gx-org/backend/shape"
)

type (
	// Node in the graph.
	Node interface {
		Graph() Graph
	}

	// Tuple bundles multiple Nodes together.
	Tuple interface {
		Node

		// Element returns a Node representing the ith element of the tuple.
		Element(i int) (Node, error)

		// Size returns the number of elements in the tuple.
		Size() int

		// Unpack returns the tuple's constituent Nodes.
		Unpack() ([]Node, error)
	}

	// Runner runs a node in a compiled graph.
	Runner interface {
		Run([]platform.Handle) (out, traces []platform.DeviceHandle, err error)
	}

	// OutputNode is an output node in the graph.
	OutputNode struct {
		Node  Node
		Shape *shape.Shape
	}

	// Graph implemented by a backend.
	// The GX interpreter uses this interface to build a graph for the backend.
	Graph interface {
		// Platform used by the graph.
		Platform() platform.Platform

		// Core returns the builder to build core operations.
		Core() CoreBuilder

		// Num returns the implementation for functions in the num package.
		Num() NumBuilder

		// Math returns the implementation for functions in the math package.
		Math() MathBuilder

		// Compile the graph for a given device.
		// The graph is not supposed to be modified once it has been compiled.
		Compile(dev platform.Device, output, traced []*OutputNode, params []*shape.Shape) (Runner, error)
	}

	// Subgraph bundles a Graph and its output node together.
	Subgraph struct {
		Graph  Graph
		Result OutputNode
	}

	// CoreBuilder creates node in the graph for core operations.
	CoreBuilder interface {
		// Graph returns the graph in which the nodes are created into.
		Graph() Graph

		// Constant returns a node representing a numerical constant value in the graph.
		Constant(value platform.HostBuffer) (Node, error)

		// Tuple returns a node representing a tuple of nodes.
		Tuple(nodes []Node) (Tuple, error)

		// Call returns a node that invokes a subgraph.
		Call(sg *Subgraph, args ...Node) (Node, error)

		// Subgraph returns a Graph instance that maps to a new subgraph.
		Subgraph(name string) (Graph, error)

		// Argument returns a node set by a caller when calling the function.
		Argument(name string, shape *shape.Shape, index int) (Node, error)

		// Unary returns a node applying a unary operator to a node.
		Unary(op *ast.UnaryExpr, x Node) (Node, error)

		// Binary returns a node applying a binary operator between two nodes.
		Binary(op *ast.BinaryExpr, x, y Node) (Node, error)

		// Reshape returns a reshape operator node.
		Reshape(x Node, axisLengths []int) (Node, error)

		// Concat concatenates multiple arrays into a single array.
		Concat(axis int, nodes []Node) (Node, error)

		// Cast returns a cast/convert operator node.
		Cast(x Node, target dtype.DataType) (Node, error)

		// Slice returns a slice on a node.
		Slice(x Node, index int) (Node, error)

		// Set returns a node to set a slice in an array.
		Set(x, updates, index Node) (Node, error)

		// DotGeneral returns a general dot operator node.
		DotGeneral(x, y Node, batchAxes, reduceAxes [2][]int) (Node, error)

		// While returns a while loop node.
		While(cond, body *Subgraph, state Node) (Node, error)

		// BroadcastInDim broadcasts data across a given set of axis.
		BroadcastInDim(x Node, shape *shape.Shape, broadcastAxes []int) (Node, error)
	}

	// NumBuilder creates node in the graph for functions in the num package from the standard library.
	NumBuilder interface {
		// Iota returns a node filling an array with values from 0 to number of elements-1.
		Iota(sh *shape.Shape, iotaAxis int) (Node, error)
	}

	// MathBuilder creates node in the graph for functions in the max package from the standard library.
	MathBuilder interface {
		// Cos returns the computation for cosine.
		Cos(x Node) (Node, error)
		// Sin returns the computation for sine.
		Sin(x Node) (Node, error)
		// Tanh returns the computation for hyperbolic tangent.
		Tanh(x Node) (Node, error)
	}
)

// String representation of an output node.
func (out *OutputNode) String() string {
	return fmt.Sprintf("%s: %v", out.Shape.String(), out.Node)
}
