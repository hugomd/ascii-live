package frames

import (
	"fmt"
	"math"
	"time"
)

var AnalogClock = FrameType{
	GetFrame:  getAnalogClockFrame,
	GetLength: getAnalogClockLength,
	GetSleep:  DefaultGetSleep(), // use default sleep for smooth animation
}

// Parameters for the clock display
const (
	width  = 60
	height = 30
	rx     = 25.0
	ry     = 12.0
	cx     = width / 2
	cy     = height / 2
)

// getAnalogClockFrame returns the current frame of the clock as a string
func getAnalogClockFrame(_ int) string {
	// Create grid
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	now := time.Now()

	// Minute markers
	for angle := 0; angle < 360; angle += 6 {
		if angle%30 == 0 {
			continue
		}
		rad := (float64(angle) - 90) * math.Pi / 180
		x := int(float64(cx) + rx*math.Cos(rad))
		y := int(float64(cy) + ry*math.Sin(rad))
		if x >= 0 && x < width && y >= 0 && y < height {
			grid[y][x] = '.'
		}
	}

	// Clock numbers
	for i := 1; i <= 12; i++ {
		angle := float64(i*30 - 90)
		rad := angle * math.Pi / 180
		x := int(float64(cx) + rx*math.Cos(rad))
		y := int(float64(cy) + ry*math.Sin(rad))
		numStr := fmt.Sprintf("%d", i)
		for j, ch := range numStr {
			posX := x - len(numStr)/2 + j
			if posX >= 0 && posX < width && y >= 0 && y < height {
				grid[y][posX] = ch
			}
		}
	}

	// Function to draw hands
	drawHand := func(lengthRatio float64, angle float64, char rune) {
		rad := (angle - 90) * math.Pi / 180
		steps := math.Max(rx, ry)
		targetRx := rx * lengthRatio
		targetRy := ry * lengthRatio
		for step := 1; step <= int(steps*lengthRatio); step++ {
			dist := float64(step) / (steps * lengthRatio)
			lx := int(float64(cx) + targetRx*dist*math.Cos(rad))
			ly := int(float64(cy) + targetRy*dist*math.Sin(rad))
			if lx >= 0 && lx < width && ly >= 0 && ly < height {
				if grid[ly][lx] < '0' || grid[ly][lx] > '9' {
					grid[ly][lx] = char
				}
			}
		}
	}

	// Calculate angles
	sAngle := float64(now.Second()) * 6
	mAngle := float64(now.Minute())*6 + float64(now.Second())*0.1
	hAngle := float64(now.Hour()%12)*30 + float64(now.Minute())*0.5

	// Draw hands
	drawHand(0.45, hAngle, '#')
	drawHand(0.75, mAngle, '*')
	drawHand(0.90, sAngle, '+')

	// Center
	grid[cy][cx] = 'O'

	// Build the frame string
	frame := "\n"
	for _, row := range grid {
		frame += string(row) + "\n"
	}

	// Footer like your original clock
	frame += fmt.Sprintf("\n%19s<JCPL_CHXD17>\n", "")
	frame += fmt.Sprintf("\n%19sDigital Time: %s\n", "", now.Format("15:04:05"))
	frame += fmt.Sprintf("%20sPress Ctrl+C to exit\n", "")

	return frame
}

// getAnalogClockLength returns 0 for infinite frames
func getAnalogClockLength() int {
	return 0
}
