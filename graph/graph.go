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

// Package graph defines the graph for a backend.
package graph

import (
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
	// The GX interpreter ses this interface to build a graph for the backend.
	Graph interface {
		// Platform used by the graph.
		Platform() platform.Platform

		// Core returns the builder to build core operations.
		Core() CoreBuilder

		// Compile the graph for a given device.
		// The graph is not supposed to be modified once it has been compiled.
		Compile(dev platform.Device, output, traced []*OutputNode, params []*shape.Shape) (Runner, error)
	}

	// CoreBuilder creates node in the graph for core operations.
	CoreBuilder interface {
		// Graph returns the graph in which the nodes are created into.
		Graph() Graph

		// NewConstant returns a node representing a numerical constant value in the graph.
		NewConstant(value platform.HostBuffer) (Node, error)

		// NewCall returns a node that invokes a subgraph with the given result node.
		NewCall(sg Graph, resultNode Node) (Node, error)

		// NewSubgraph returns a Graph instance that maps to a new subgraph.
		NewSubgraph(name string) (Graph, error)

		// NewArgument returns a node set by a caller when calling the function.
		NewArgument(name string, shape *shape.Shape, index int) (Node, error)

		// NewUnary returns a node applying a unary operator to a node.
		NewUnary(op *ast.UnaryExpr, x Node) (Node, error)

		// NewBinary returns a node applying a binary operator between two nodes.
		NewBinary(op *ast.BinaryExpr, x, y Node) (Node, error)

		// NewReshape returns a reshape operator node.
		NewReshape(x Node, axisLengths []int) (Node, error)

		// NewConcat concatenates multiple arrays into a single array.
		NewConcat(axis int, nodes []Node) (Node, error)

		// NewCast returns a cast/convert operator node.
		NewCast(x Node, target dtype.DataType) (Node, error)

		// NewSlice returns a slice on a node.
		NewSlice(x Node, index int) (Node, error)

		// NewSet returns a node to set a slice in an array.
		NewSet(x, updates, index Node) (Node, error)

		// NewDotGeneral returns a general dot operator node.
		NewDotGeneral(x, y Node, batchAxes, reduceAxes [2][]int) (Node, error)
	}
)
