// Package backend defines the abstraction to implement for a GX backend.
package backend

import (
	"github.com/gx-org/backend/graph"
	"github.com/gx-org/backend/platform"
)

// Backend is a GX backend.
type Backend interface {
	// Platform supporting the backend.
	Platform() platform.Platform

	// NewGraph returns a new graph backend.
	NewGraph(name string) graph.Graph
}
