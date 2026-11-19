package frames

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadCustomFrames(t *testing.T) {
	dir := t.TempDir()
	set := filepath.Join(dir, "custom-test")
	if err := os.Mkdir(set, 0755); err != nil {
		t.Fatal(err)
	}
	for name, data := range map[string]string{
		"delay.env": "45\n", "frame-001.ansi": "second", "frame-000.ansi": "\x1b[31mfirst",
	} {
		if err := os.WriteFile(filepath.Join(set, name), []byte(data), 0644); err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() { delete(FrameMap, "custom-test") })
	if err := LoadCustomFrames(dir); err != nil {
		t.Fatal(err)
	}
	animation := FrameMap["custom-test"]
	if animation.GetLength() != 2 || animation.GetFrame(0) != "\x1b[31mfirst" || animation.GetFrame(1) != "second" || animation.GetSleep() != 45*time.Millisecond {
		t.Fatal("custom frame content, order, or delay was not preserved")
	}
}

func TestLoadCustomFramesRejectsInvalidDelay(t *testing.T) {
	dir := t.TempDir()
	set := filepath.Join(dir, "bad")
	if err := os.Mkdir(set, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(set, "delay.env"), []byte("0"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := LoadCustomFrames(dir); err == nil {
		t.Fatal("expected invalid delay to fail")
	}
}
