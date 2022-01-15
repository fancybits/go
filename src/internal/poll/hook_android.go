// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build android
// +build android

package poll

import "syscall"

// CloseFunc is used to hook the close call.
func CloseFunc(fd int) error {
	err := syscall.Close(fd)
	if err == syscall.EPERM {
		return nil
	}
	return err
}

// AcceptFunc is used to hook the accept call.
var AcceptFunc func(int) (int, syscall.Sockaddr, error) = syscall.Accept
