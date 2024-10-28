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

// Package platform defines the interface for a GX
//
// A GX platform manages a host and its associated devices,
// also providing the means to exchange data among them.
package platform

import "github.com/gx-org/backend/shape"

type (
	// Platform is a host orchestrating one or more devices.
	Platform interface {
		// Name of the platform.
		Name() string

		// Device returns the device managed by the backend.
		Device(int) (Device, error)
	}

	// Device running GX code.
	Device interface {
		// Platform returns the host and its devices owning this device.
		Platform() Platform

		// Send raw data to the device.
		Send(buf []byte, sh *shape.Shape) (DeviceHandle, error)
	}
)
