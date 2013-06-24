// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by a BSD-style (2-Clause)
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/errnoh/go.tbi"
)

func main() {
	w, h, c := tbi.Start()
	buf := tbi.NewBuffer(w, h, nil)
	printcolors(buf)
	buf.Draw()
	<-c
	tbi.Quit()
}

func printcolors(b *tbi.Buffer) {
	var (
		tmp   byte
		start byte
	)

        width, _ := b.Dimensions()
        runes, _, bg := b.Data()

	for i := 0; i < len(runes); i++ {
		runes[i] = ' '
		bg[i] = tmp

		switch {
		case tmp == 255:
			copy(runes[i+3:], []rune(fmt.Sprintf(" %03d-%03d", start, tmp)))
			return
		case tmp == 15 || (tmp > 15 && tmp%6 == 3):
			copy(runes[i+3:], []rune(fmt.Sprintf(" %03d-%03d", start, tmp)))
			i += width - (i % width)
			start = tmp + 1
		}

		tmp++
	}
}
