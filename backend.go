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

// Package backend defines the abstraction to implement for a GX backend.
package backend

import (
	"github.com/gx-org/backend/shapes"
)

// Backend is a GX backend.
type Backend interface {
	// Platform supporting the backend.
	Platform() Platform

	// Builder returns a new ops builder.
	Builder(name string) (Builder, error)

	// Finalize everything linked to the backend.
	// It is invalid to use the platform or graph builder after this call.
	Finalize() error
}

// Builder builds a computation.
type Builder interface {
	// Name of the computation being built.
	Name() string

	// Main returns the main function of this computation, named MainName.
	// Operations added to Main become part of the compiled computation.
	// This is the default function where all operations should be added
	// unless explicitly building a sub-function.
	Main() Function

	// Compile the computation built for a given device.
	// The graph is not supposed to be modified once it has been compiled.
	// This immediately invalidates the Builder and returns an Executable
	// that can be used to run the computation.
	Compile(dev DeviceNum, output, traced []*OutputNode, params []*shapes.Shape) (Executable, error)
}
