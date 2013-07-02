// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by a BSD-style (2-Clause)
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/errnoh/color"
	"github.com/errnoh/go.tbi"
	"image/png"
	"log"
	"os"
)

var usecolor = flag.Bool("color", true, "true for 256 colors, false for greyscale")

func main() {
	flag.Parse()
	w, h, c := tbi.Start()
	buf := tbi.NewBuffer(w, h, nil)
	drawcircle(buf, *usecolor)
	buf.Draw()
	<-c
	tbi.Quit()
}

func drawcircle(buf *tbi.Buffer, colors bool) {
	f, err := os.Open("colorwheel.png")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Println(err)
		return
	}

	r := img.Bounds()
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			c := img.At(x, y)
			var val byte
			if colors {
                            val = color.Term256Model.Convert(c).(color.Term256).Val
			} else {
                            val = color.Term256GreyscaleModel.Convert(c).(color.Term256).Val
			}
			// font height:width is quite close to 2:1 so fill 2 columns for each row
			buf.Set(2*x, y, ' ', 0, val)
			buf.Set(2*x+1, y, ' ', 0, val)
		}
	}
}
