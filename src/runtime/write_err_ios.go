// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

var (
	WriteErrFD uintptr = 2
)

func writeErr(b []byte) {
	write(WriteErrFD, unsafe.Pointer(&b[0]), int32(len(b)))
}
