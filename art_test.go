package art

import (
	"image"
	"image/color"
	"strings"
	"testing"
)

func TestFromImageResizesAndPreservesTerminalShape(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 160, 80))
	for y := 0; y < 80; y++ {
		for x := 0; x < 160; x++ {
			if x < 80 {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	got, err := FromImage(img, Options{MaxWidth: 40})
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(got, "\n")
	if len(lines) != 10 {
		t.Fatalf("expected 10 lines, got %d", len(lines))
	}
	if len(lines[0]) != 40 {
		t.Fatalf("expected 40 columns, got %d", len(lines[0]))
	}
	if lines[0][0] != '@' || lines[0][39] != ' ' {
		t.Fatalf("unexpected luminance mapping: %q", lines[0])
	}
}

func TestFromImageRejectsNil(t *testing.T) {
	if _, err := FromImage(nil, Options{}); err == nil {
		t.Fatal("expected nil image error")
	}
}
