package frames

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// LoadCustomFrames overlays generated ANSI animations on the built-in ones.
func LoadCustomFrames(dir string) error {
	sets, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, set := range sets {
		if !set.IsDir() {
			continue
		}
		path := filepath.Join(dir, set.Name())
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		var content []string
		var delay time.Duration
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			data, err := os.ReadFile(filepath.Join(path, entry.Name()))
			if err != nil {
				return err
			}
			if entry.Name() == "delay.env" {
				ms, err := strconv.Atoi(strings.TrimSpace(string(data)))
				if err != nil || ms <= 0 {
					return fmt.Errorf("invalid delay in %s", path)
				}
				delay = time.Duration(ms) * time.Millisecond
			} else if strings.HasSuffix(entry.Name(), ".ansi") {
				content = append(content, string(data))
			}
		}
		if len(content) == 0 || delay == 0 {
			return fmt.Errorf("missing frames or delay in %s", path)
		}
		animation := DefaultFrameType(content)
		animation.GetSleep = func() time.Duration { return delay }
		FrameMap[set.Name()] = animation
	}
	return nil
}
