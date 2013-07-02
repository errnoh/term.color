// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by a BSD-style (2-Clause)
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
	"testing"
)

type colorTest struct {
	val  uint8
	rgba color.RGBA
}

func TestToRGBA(t *testing.T) {
	var toRGBATests = []colorTest{
		{0, color.RGBA{0, 0, 0, 0}}, {1, color.RGBA{187, 0, 0, 0}}, {2, color.RGBA{0, 187, 0, 0}}, {7, color.RGBA{187, 187, 187, 0}},
		{8, color.RGBA{85, 85, 85, 0}}, {9, color.RGBA{255, 85, 85, 0}}, {10, color.RGBA{85, 255, 85, 0}}, {15, color.RGBA{255, 255, 255, 0}},
		{16, color.RGBA{0, 0, 0, 0}}, {17, color.RGBA{0, 0, 95, 0}}, {18, color.RGBA{0, 0, 135, 0}}, {24, color.RGBA{0, 95, 135, 0}}, {231, color.RGBA{255, 255, 255, 0}},
		{232, color.RGBA{8, 8, 8, 0}}, {233, color.RGBA{18, 18, 18, 0}}, {255, color.RGBA{238, 238, 238, 0}},
	}

	for _, pair := range toRGBATests {
		expected := pair.rgba
		result := color.RGBAModel.Convert(Term256{pair.val}).(color.RGBA)

		if result.R != expected.R || result.G != expected.G || result.B != expected.B || result.A != expected.A {
			t.Fatalf("Converted %d, expected %#v, got %#v", pair.val, expected, result)
		}
	}
}
