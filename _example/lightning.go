// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by a BSD-style (2-Clause)
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/errnoh/color"
	"github.com/errnoh/go.tbi"
	"image"
	"image/gif"
	"log"
	"net/http"
	"time"

	"github.com/jsummers/fpresize"
)

var usecolor = flag.Bool("color", true, "true for 256 colors, false for greyscale")

func main() {
	flag.Parse()
	w, h, c := tbi.Start()
	buf := tbi.NewBuffer(w, h, nil)

	tick := time.Tick(2 * time.Minute)
	display(buf)
loop:
	for {
		select {
		case <-c:
			break loop
		case <-tick:
			if err := display(buf); err != nil {
				continue
			}
		}
	}
	tbi.Quit()
}

func display(buf *tbi.Buffer) error {
	img, err := open("http://cdn.flcenter.net/weather/tutka1/map.jpg")
	if err != nil {
		return err
	}

	_, h := buf.Dimensions()
	r := img.Bounds()
	imgw, imgh := r.Max.X, r.Max.Y
	x, y := (imgw*h)/imgh, h

	fp := fpresize.New(img)
	fp.SetTargetBounds(image.Rect(0, 0, x*2, y))
	img, err = fp.ResizeToRGBA()
	if err != nil {
		return err
	}

	draw(buf, img, *usecolor)
	buf.Draw()
	return nil
}

func open(path string) (img image.Image, err error) {
	var (
		resp *http.Response
	)

	resp, err = http.Get(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	img, err = gif.Decode(resp.Body)

	if err != nil {
		log.Println(err)
		return
	}
	return
}

func draw(buf *tbi.Buffer, img image.Image, colors bool) {
	var val byte

	r := img.Bounds()
	for x := r.Min.X; x < r.Max.X-1; x = x + 2 {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			c := img.At(x, y)
			if colors {
                            val = color.Term256Model.Convert(c).(color.Term256).Val
			} else {
                            val = color.Term256GreyscaleModel.Convert(c).(color.Term256).Val
			}
			// font height:width is quite close to 2:1 so fill 2 columns for each row
			buf.Set(x, y, ' ', 0, val)
			buf.Set(x+1, y, ' ', 0, val)
		}
	}
}
