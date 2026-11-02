package art

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

// Options controls the terminal representation of an image.
type Options struct {
	MaxWidth int
}

// FromImage converts an image into a terminal-friendly monochrome frame.
// The height is halved because terminal characters are usually taller than wide.
func FromImage(src image.Image, opts Options) (string, error) {
	if src == nil {
		return "", fmt.Errorf("image is nil")
	}
	maxWidth := opts.MaxWidth
	if maxWidth <= 0 {
		maxWidth = 80
	}

	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	if width == 0 || height == 0 {
		return "", fmt.Errorf("image has no pixels")
	}

	outWidth := width
	if outWidth > maxWidth {
		outWidth = maxWidth
	}
	outHeight := int(float64(height) * float64(outWidth) / float64(width) * 0.5)
	if outHeight < 1 {
		outHeight = 1
	}

	const ramp = "@%#*+=-:. "
	var result strings.Builder
	result.Grow(outWidth * (outHeight + 1))

	for y := 0; y < outHeight; y++ {
		for x := 0; x < outWidth; x++ {
			sx := bounds.Min.X + x*width/outWidth
			sy := bounds.Min.Y + y*height/outHeight
			brightness := int(luminance(src.At(sx, sy)))
			result.WriteByte(ramp[brightness*(len(ramp)-1)/255])
		}
		if y < outHeight-1 {
			result.WriteByte('\n')
		}
	}
	return result.String(), nil
}

func luminance(c color.Color) uint8 {
	// RGBA returns alpha-premultiplied channel values, which already represent
	// compositing the pixel over black.
	r, g, b, _ := c.RGBA()
	return uint8((299*r + 587*g + 114*b) / 1000 / 257)
}
