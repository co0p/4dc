package icon

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "os"
)

// Generate creates a simple tomato PNG icon of given square size and writes it to outPath.
func Generate(size int, outPath string) error {
    if size <= 0 {
        return nil
    }
    img := image.NewRGBA(image.Rect(0, 0, size, size))

    clear := color.RGBA{0, 0, 0, 0}
    for y := 0; y < size; y++ {
        for x := 0; x < size; x++ {
            img.Set(x, y, clear)
        }
    }

    cx := float64(size) / 2
    cy := float64(size) / 2
    r := float64(size) * 0.38
    r2 := r * r
    red := color.RGBA{R: 220, G: 50, B: 47, A: 255}
    for y := 0; y < size; y++ {
        for x := 0; x < size; x++ {
            dx := float64(x) - cx
            dy := float64(y) - cy
            if dx*dx+dy*dy <= r2 {
                img.Set(x, y, red)
            }
        }
    }

    stemW := int(math.Max(2, float64(size)*0.08))
    stemH := int(math.Max(2, float64(size)*0.12))
    stemX0 := int(cx) - stemW/2
    stemY0 := int(cy - r - float64(stemH)*0.35)
    green := color.RGBA{R: 35, G: 130, B: 65, A: 255}
    for y := stemY0; y < stemY0+stemH; y++ {
        if y < 0 || y >= size {
            continue
        }
        for x := stemX0; x < stemX0+stemW; x++ {
            if x < 0 || x >= size {
                continue
            }
            img.Set(x, y, green)
        }
    }

    f, err := os.Create(outPath)
    if err != nil {
        return err
    }
    defer f.Close()
    return png.Encode(f, img)
}
