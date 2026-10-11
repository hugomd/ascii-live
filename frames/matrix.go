package frames

import (
	"math/rand"
	"strings"
	"time"
)

// Matrix animation
var Matrix = DefaultFrameType(matrixGenerateFrames())

type matrixDrop struct {
	pos    int
	length int
	speed  float64
}

// matrixGenerateFrames uses a fixed canvas because the server has no client terminal.
func matrixGenerateFrames() []string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	ROWS, COLS := 24, 80 // fixed curl canvas
	numFrames := 500
	const chars = "01ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%^&*"

	// initialize drops per column
	drops := make([]matrixDrop, COLS)
	for i := range drops {
		drops[i] = matrixDrop{
			pos:    rand.Intn(ROWS),
			length: 6 + rand.Intn(6),         // trail length 6-11
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
					line += string(chars[rand.Intn(len(chars))])
				} else {
					line += " "
				}
			}
			lines[r] = line
		}

		frames[f] = strings.Join(lines, "\n")

		// move each matrixDrop down
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
