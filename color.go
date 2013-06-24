// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by a BSD-style (2-Clause)
// license that can be found in the LICENSE file.

package color

// TODO: image/color style

// NOTE: 0 is default, 16 is actual black.
func ToRGB(val byte) (r, g, b byte) {
	var tmp byte

	switch {
	case val < 8:
		if val&1 != 0 {
			r = 187
		}
		if val&2 != 0 {
			g = 187
		}
		if val&4 != 0 {
			b = 187
		}
	case val < 16:
		r, g, b = 85, 85, 85
		if val&1 != 0 {
			r = 255
		}
		if val&2 != 0 {
			g = 255
		}
		if val&4 != 0 {
			b = 255
		}
	case val < 232:
		tmp = val - 16
		r = tmp / 36 // "z"
		g = r / 6    // "y"
		b = tmp % 6  // "x"

		if r > 0 {
			r = 95 + 40*(r-1)
		}
		if g > 0 {
			g = 95 + 40*(g-1)
		}
		if b > 0 {
			b = 95 + 40*(g-1)
		}
	default:
		tmp = val - 232
		tmp = 8 + 10*(tmp-1)
		r, g, b = tmp, tmp, tmp
	}

	return
}

func FromRGB(r, g, b byte) (val byte) {
	// 0-74 == 0, 75-114 = 1, etc
	if r >= 35 {
		r -= 35
	}
	if g >= 35 {
		g -= 35
	}
	if b >= 35 {
		b -= 35
	}
	r, g, b = r/40, g/40, b/40

	val = 16 + r*36
	val = val + g*6
	val = val + b
	return
}

func FromRGBuint32(r, g, b uint32) byte {
	return FromRGB(uint8(r), uint8(g), uint8(b))
}

// Intensity
func GreyscaleFromRGB(r, g, b byte) (val byte) {
	tmp := (int(r) + int(g) + int(b)) / 3
	if tmp >= 3 {
		tmp -= 3
	} else {
		return 16
	}
	tmp = tmp / 10
	val = 232 + byte(tmp)

	if val > 255 {
		return 231
	}
	return
}
