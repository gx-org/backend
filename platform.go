// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import "github.com/gx-org/backend/shapes"

type (
	// Platform is a host orchestrating one or more devices.
	Platform interface {
		// Name of the platform.
		Name() string

		// Send raw data to the device.
		Send(dev DeviceNum, buf []byte, sh *shapes.Shape) (DeviceHandle, error)

		// Finalize everything linked to the platform.
		// It is invalid to use any device from the platform after this call.
		Finalize() error
	}

	// DeviceNum represents a device ordinal on the platform.
	DeviceNum int
)
