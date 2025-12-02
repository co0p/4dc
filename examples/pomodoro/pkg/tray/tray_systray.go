package tray

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"

	"github.com/getlantern/systray"
)

// SystrayAdapter shows a tray icon and menu using getlantern/systray.
type SystrayAdapter struct{}

func (a SystrayAdapter) Show(title string, items []Item) {
	systray.Run(func() {
		log.Println("[tray] systray starting")

		if title != "" {
			// Icon-only: do not set a title to avoid text in the menu bar
			systray.SetTooltip(title)
		}

		// Use a high-contrast monochrome template icon so macOS auto-inverts for the menu bar
		systray.SetTemplateIcon(generateTemplateGlyph(), generateTemplateGlyph())

		for _, it := range items {
			m := systray.AddMenuItem(it.Title, "")
			go func(h func()) {
				for range m.ClickedCh {
					if h != nil {
						h()
					}
				}
			}(it.OnClick)
		}

		quit := systray.AddMenuItem("Quit", "exit app")
		go func() {
			for range quit.ClickedCh {
				log.Println("[tray] quit clicked, exiting")
				systray.Quit()
			}
		}()
	}, func() {
		log.Println("[tray] systray exiting")
	})
}

// generateTemplateGlyph creates an 18x18 monochrome PNG suitable as a macOS template icon.
// It draws a bold circular glyph with a small stem to resemble a pomodoro tomato.
func generateTemplateGlyph() []byte {
	const (
		w = 18
		h = 18
	)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	// Transparent background
	clear := color.RGBA{0, 0, 0, 0}
	white := color.RGBA{255, 255, 255, 255}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, clear)
		}
	}
	// Bold filled circle
	cx, cy := 9.0, 9.0
	r := 7.0
	r2 := r * r
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := float64(x) - cx
			dy := float64(y) - cy
			if dx*dx+dy*dy <= r2 {
				img.Set(x, y, white)
			}
		}
	}
	// Small stem at the top center
	for y := 2; y <= 4; y++ {
		for x := 8; x <= 10; x++ {
			img.Set(x, y, white)
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		log.Printf("[tray] failed to encode template icon png: %v\n", err)
		return nil
	}
	return buf.Bytes()
}
