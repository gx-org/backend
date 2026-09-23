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
	"github.com/gx-org/backend/shapes"
)

type (
	// Handle to an array managed by the platform.
	Handle interface {
		// Shape of the underlying array.
		Shape() *shapes.Shape

		// ToDevice transfers the handle to a device.
		ToDevice(DeviceNum) (DeviceHandle, error)

		// ToHost fetches the data from the handle and write it to buffer.
		ToHost(buffer []byte) error
	}

	// DeviceHandle is an array located on a device.
	DeviceHandle interface {
		Handle

		// Device on which the array is located.
		Device() DeviceNum
	}
)
