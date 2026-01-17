package frames

import (
	"math/rand"
	"strings"
	"time"

	"github.com/buger/goterm" // only for terminal size detection
)

// Matrix animation
var Matrix = DefaultFrameType(generateFrames())

type drop struct {
	pos    int
	length int
	speed  float64
}

// generateFrames dynamically generates frames using terminal size
func generateFrames() []string {
	rand.Seed(time.Now().UnixNano())

	ROWS, COLS := goterm.Height(), goterm.Width() // get terminal size dynamically
	numFrames := 500

	// initialize drops per column
	drops := make([]drop, COLS)
	for i := range drops {
		drops[i] = drop{
			pos:    rand.Intn(ROWS),
			length: 6 + rand.Intn(6),      // trail length 6-11
			speed:  0.8 + rand.Float64()*0.2, // chance to move per frame
		}
	}

	frames := make([]string, numFrames)

	for f := 0; f < numFrames; f++ {
		lines := make([]string, ROWS)
		for r := 0; r < ROWS; r++ {
			line := ""
			for c := 0; c < COLS; c++ {
				d := drops[c]
				if r <= d.pos && r > d.pos-d.length {
					line += string(randomChar())
				} else {
					line += " "
				}
			}
			lines[r] = line
		}

		frames[f] = strings.Join(lines, "\n")

		// move each drop down
		for c := range drops {
			if rand.Float64() < drops[c].speed {
				drops[c].pos++
				if drops[c].pos >= ROWS+drops[c].length {
					drops[c].pos = 0
					drops[c].length = 6 + rand.Intn(6)
				}
			}
		}
	}

	return frames
}

func randomChar() rune {
	chars := "01ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%^&*"
	return rune(chars[rand.Intn(len(chars))])
}
