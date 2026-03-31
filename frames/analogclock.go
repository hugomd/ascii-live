package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	width := 60
	height := 30
	cx, cy := width/2, height/2

	rx := 25.0
	ry := 12.0

	for {
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

		// Numbers
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

		// Draw hand function
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

		// Angles
		sAngle := float64(now.Second()) * 6
		mAngle := float64(now.Minute())*6 + float64(now.Second())*0.1
		hAngle := float64(now.Hour()%12)*30 + float64(now.Minute())*0.5

		// Draw hands
		drawHand(0.45, hAngle, '#')
		drawHand(0.75, mAngle, '*')
		drawHand(0.90, sAngle, '+')

		grid[cy][cx] = 'O'

		// Clear screen (ANSI - ascii.live friendly)
		fmt.Print("\033[H\033[2J")

		fmt.Println()
		for _, row := range grid {
			fmt.Println(string(row))
		}

		fmt.Printf("\n%19s<JCPL_CHXD17>\n", "")
		fmt.Printf("\n%19sDigital Time: %s\n", "", now.Format("15:04:05"))
		fmt.Printf("%20sPress Ctrl+C to exit\n", "")

		time.Sleep(50 * time.Millisecond)
	}
}
