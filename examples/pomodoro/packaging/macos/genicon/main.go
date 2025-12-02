package main

import (
"flag"
"image"
"image/color"
"image/png"
"log"
"math"
"os"
)

func main() {
	size := flag.Int("size", 1024, "square icon size in pixels")
	out := flag.String("out", "icon.png", "output PNG path")
	flag.Parse()

	if *size <= 0 {
		log.Fatal("size must be > 0")
	}

	img := image.NewRGBA(image.Rect(0, 0, *size, *size))

	// Transparent background
	clear := color.RGBA{0, 0, 0, 0}
	for y := 0; y < *size; y++ {
		for x := 0; x < *size; x++ {
			img.Set(x, y, clear)
		}
	}

	// Tomato body: red filled circle
	cx := float64(*size) / 2
	cy := float64(*size) / 2
	r := float64(*size) * 0.38 // radius ~38% of size
	r2 := r * r
	red := color.RGBA{R: 220, G: 50, B: 47, A: 255}
	for y := 0; y < *size; y++ {
		for x := 0; x < *size; x++ {
			dx := float64(x) - cx
			dy := float64(y) - cy
			if dx*dx+dy*dy <= r2 {
				img.Set(x, y, red)
			}
		}
	}

	// Tomato stem: green rectangle at top center
	stemW := int(math.Max(2, float64(*size)*0.08))
	stemH := int(math.Max(2, float64(*size)*0.12))
	stemX0 := int(cx) - stemW/2
	stemY0 := int(cy - r - float64(stemH)*0.35)
	green := color.RGBA{R: 35, G: 130, B: 65, A: 255}
	for y := stemY0; y < stemY0+stemH; y++ {
		if y < 0 || y >= *size {
			continue
		}
		for x := stemX0; x < stemX0+stemW; x++ {
			if x < 0 || x >= *size {
				continue
			}
			img.Set(x, y, green)
		}
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
