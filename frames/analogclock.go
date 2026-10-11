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
	analogClockWidth  = 60
	analogClockHeight = 30
	analogClockRx     = 25.0
	analogClockRy     = 12.0
	analogClockCx     = analogClockWidth / 2
	analogClockCy     = analogClockHeight / 2
)

// getAnalogClockFrame returns the current frame of the clock as a string
func getAnalogClockFrame(_ int) string {
	// Create grid
	grid := make([][]rune, analogClockHeight)
	for i := range grid {
		grid[i] = make([]rune, analogClockWidth)
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
		x := int(float64(analogClockCx) + analogClockRx*math.Cos(rad))
		y := int(float64(analogClockCy) + analogClockRy*math.Sin(rad))
		if x >= 0 && x < analogClockWidth && y >= 0 && y < analogClockHeight {
			grid[y][x] = '.'
		}
	}

	// Clock numbers
	for i := 1; i <= 12; i++ {
		angle := float64(i*30 - 90)
		rad := angle * math.Pi / 180
		x := int(float64(analogClockCx) + analogClockRx*math.Cos(rad))
		y := int(float64(analogClockCy) + analogClockRy*math.Sin(rad))
		numStr := fmt.Sprintf("%d", i)
		for j, ch := range numStr {
			posX := x - len(numStr)/2 + j
			if posX >= 0 && posX < analogClockWidth && y >= 0 && y < analogClockHeight {
				grid[y][posX] = ch
			}
		}
	}

	// Function to draw hands
	drawHand := func(lengthRatio float64, angle float64, char rune) {
		rad := (angle - 90) * math.Pi / 180
		steps := math.Max(analogClockRx, analogClockRy)
		targetRx := analogClockRx * lengthRatio
		targetRy := analogClockRy * lengthRatio
		for step := 1; step <= int(steps*lengthRatio); step++ {
			dist := float64(step) / (steps * lengthRatio)
			lx := int(float64(analogClockCx) + targetRx*dist*math.Cos(rad))
			ly := int(float64(analogClockCy) + targetRy*dist*math.Sin(rad))
			if lx >= 0 && lx < analogClockWidth && ly >= 0 && ly < analogClockHeight {
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
	grid[analogClockCy][analogClockCx] = 'O'

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
