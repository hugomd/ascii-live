package frames_test

import (
	"testing"

	"github.com/hugomd/ascii-live/frames"
)

func TestRegisteredFrames(t *testing.T) {
	for name, source := range frames.FrameMap {
		t.Run(name, func(t *testing.T) {
			if source.GetSleep() <= 0 {
				t.Fatal("frame delay must be positive")
			}
			length := source.GetLength()
			if length < 0 {
				t.Fatal("frame length must not be negative")
			}
			if first := source.GetFrame(0); first == "" {
				t.Fatal("first frame must not be empty")
			} else if length > 0 && source.GetFrame(length) != first {
				t.Fatal("animation must wrap to its first frame")
			}
			for i := 0; i < length; i++ {
				if source.GetFrame(i) == "" {
					t.Fatalf("frame %d is empty", i)
				}
			}
		})
	}
}
