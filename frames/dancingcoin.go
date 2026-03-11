package main

import (
	"fmt"
	"time"
)

func main() {
	// 10 Full-Height, High-Density Frames
	// Each frame is a complete "Bajillion-Character" block
	frames := []string{
		// FRAME 1
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
 ffffffffttft
 ffffffffttff
 fffffffttttt`,

		// FRAME 2
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
 ffffffffttft
  ffffffffttf
   fffffffttt`,

		// FRAME 3
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
  ffffffffttf
   fffffffftt
    ffffffftt`,

		// FRAME 4
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
   fffffffftt
    fffffffft
     fffffftt`,

		// FRAME 5
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
    fffffffft
     ffffffff
      fffffft`,

		// FRAME 6
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
     ffffffft
      fffffff
       ffffff`,

		// FRAME 7
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
    fffffffft
     ffffffff
      fffffft`,

		// FRAME 8
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
   fffffffftt
    fffffffft
     fffffftt`,

		// FRAME 9
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
  ffffffffttf
   fffffffftt
    ffffffftt`,

		// FRAME 10
		`      .---.
     / | \ \
    /  |  \ \
   |   |   | |
    \  |  / /
     '---'
    LLfftfLL
   fttttffff
  ttttttttttt
 ffffffffttft
  ffffffffttf
   fffffffttt`,
	}

	// Hide cursor and clear screen
	fmt.Print("\033[?25l\033[2J")

	for {
		for i, frame := range frames {
			// Reset to top-left
			fmt.Print("\033[H")

			// Color cycling: Alternate between Cyan and Magenta
			if i%2 == 0 {
				fmt.Print("\033[36m") // Cyan
			} else {
				fmt.Print("\033[35m") // Magenta
			}

			fmt.Println(frame)
			fmt.Print("\033[0m") // Reset color

			// 100ms Sleep as requested
			time.Sleep(100 * time.Millisecond)
		}
	}
}
