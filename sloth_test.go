// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/clivern/sloth/pkg"
)

// TestPkgName test cases
func TestPkgName(t *testing.T) {
	pkg.Expect(t, "Sloth", "Sloth")
}
