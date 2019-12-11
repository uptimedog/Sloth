// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"os"
)

// FS struct
type FS struct{}

// PathExists reports whether the path exists
func (fs *FS) PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileExists reports whether the named file exists
func (fs *FS) FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// DirExists reports whether the dir exists
func (fs *FS) DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

// EnsureDir ensures that directory exists
func (fs *FS) EnsureDir(dirName string, mode int) (bool, error) {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return true, nil
	}
	return false, err
}
