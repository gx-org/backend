package platform

import (
	"github.com/pkg/errors"
	"github.com/gx-org/backend/shape"
)

type (
	// Handle to an array managed by the platform.
	Handle interface {
		// Shape of the underlying array.
		Shape() *shape.Shape

		// ToDevice transfers the handle to a device.
		ToDevice(Device) (DeviceHandle, error)

		// ToHost fetches the data from the handle and write it to buffer.
		ToHost(buffer HostBuffer) error
	}

	// DeviceHandle is an array located on a device.
	DeviceHandle interface {
		Handle

		// Device on which the array is located.
		Device() Device
	}

	// HostBuffer is a handle to a buffer of data located locally on the platform,
	// typically in CPU memory, and shared between a platform and its users.
	//
	// The data needs to be acquired before being written or read. All access to the data
	// is blocked until that is released.
	HostBuffer interface {
		Handle
		// Acquire locks the buffer and returns it.
		// The buffer can be read or written by the caller. All other access is locked.
		// Returns nil if the handle has been freed.
		Acquire() []byte
		// Release the buffer. The caller of that function should not read or write data
		// from the buffer.
		Release()
		// Free the memory occupied by the buffer. The handle is invalid after calling this function.
		Free()
	}

	// Allocator allocates memory on the host.
	Allocator interface {
		Allocate(*shape.Shape) (HostBuffer, error)
	}
)

// HostTransfer transfers data from a source host buffer to another.
func HostTransfer(dstB, srcB HostBuffer) error {
	src := srcB.Acquire()
	defer srcB.Release()
	dst := dstB.Acquire()
	defer dstB.Release()
	if len(src) != len(dst) {
		return errors.Errorf("cannot transfer data from a buffer of size %d to a buffer of size %d", len(src), len(dst))
	}
	copy(src, dst)
	return nil
}
