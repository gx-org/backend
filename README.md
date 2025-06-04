# Repository backend: GX interface specification

Repository backend defines a set of interface to implement to provide
GX with a new backend.

A backend is composed of:

* a platform: a platform provides one or more devices (i.e. accelerators)
and the methods to transfer data to and from a device.
* a graph builder: primitives to build a compute graph for device. A compute
graph will be built by GX, compile for a given device, and then run from
a host language.

## Disclaimer

This is not an official Google DeepMind product (experimental or otherwise), it is
just code that happens to be owned by Google-DeepMind.

As of today, we do not consider any part of the language as stable. Breaking
changes will happen on a regular basis.

You are welcome to send PR or to report bugs. We will do our best to answer but there
is no guarantee that you will get a response.
