package frames

import (
	"strings"
	"time"
)

const (
	duckRaceWidth  = 77
	duckRaceHeight = 27
	duckRaceLength = 122
	duckFinish     = 66

	duckRaceReset = "\x1b[0m"
	duckRaceWater = "\x1b[36m"
	duckRaceGold  = "\x1b[33m"
	duckRaceBill  = "\x1b[93m"
)

// DuckRace follows three ducks through a sprint and a mid-race snack break.
var DuckRace = FrameType{
	GetFrame:  duckRaceFrame,
	GetLength: func() int { return duckRaceLength },
	GetSleep:  func() time.Duration { return 90 * time.Millisecond },
}

// Each keyframe gives the column of each duck's tail. Their bills extend
// 16 cells to the right. The finish is decided by the tip of the bill.
var duckRacePositions = []struct {
	tick int
	x    [3]int
}{
	{0, [3]int{11, 11, 11}},
	{17, [3]int{11, 11, 11}},
	{32, [3]int{23, 18, 17}},
	{44, [3]int{28, 29, 22}},
	{54, [3]int{32, 35, 23}},
	{66, [3]int{38, 40, 23}},
	{79, [3]int{44, 43, 34}},
	{90, [3]int{48, 47, 45}},
	{94, [3]int{49, 48, 50}},
	{96, [3]int{50, 49, 51}},
	{98, [3]int{51, 50, 52}},
	{104, [3]int{53, 52, 55}},
}

func duckRacePosition(tick, lane int) int {
	prev := duckRacePositions[0]
	for _, next := range duckRacePositions[1:] {
		if tick <= next.tick {
			return prev.x[lane] + (next.x[lane]-prev.x[lane])*(tick-prev.tick)/(next.tick-prev.tick)
		}
		prev = next
	}
	return prev.x[lane]
}

func duckRaceCall(tick int) string {
	switch {
	case tick < 6:
		return "On your marks ... 3"
	case tick < 12:
		return "On your marks ... 2"
	case tick < 18:
		return "On your marks ... 1"
	case tick < 31:
		return "QUACK! Captain takes the lead."
	case tick < 44:
		return "Here comes Punk ..."
	case tick < 55:
		return "Punk noses ahead!"
	case tick < 68:
		return "Wait. Is Crumb stopping for a snack?"
	case tick < 82:
		return "Now THAT was fast food!"
	case tick < 94:
		return "Three ducks, neck and neck ..."
	default:
		return "CRUMB WINS! By a breadcrumb."
	}
}

func duckRaceFrame(i int) string {
	tick := i % duckRaceLength
	var pond [duckRaceHeight][duckRaceWidth]byte
	var colors [duckRaceHeight][duckRaceWidth]string
	for y := range pond {
		for x := range pond[y] {
			pond[y][x] = ' '
		}
	}
	// Store color separately so escape sequences never use a cell on the track.
	put := func(x, y int, s, color string) {
		copy(pond[y][x:], s)
		for k := range s {
			colors[y][x+k] = color
		}
	}
	put(1, 0, "QUACK GRAND PRIX", duckRaceWater)
	put(55, 0, "DUCK POND   * LIVE", duckRaceWater)
	put(0, 1, strings.Repeat("=", duckRaceWidth), "")
	put(11, 2, "START", "")
	put(duckFinish-2, 2, "FINISH", "")

	for lane, name := range []string{"CAPTAIN", "PUNK", "CRUMB"} {
		y := 3 + lane*7
		laneColor := []string{"\x1b[32m", "\x1b[31m", duckRaceGold}[lane]
		put(1, y+1, []string{"[01]", "[02]", "[03]"}[lane], laneColor)
		put(1, y+2, name, laneColor)
		for row := 0; row < 6; row++ {
			put(9, y+row, "|", "")
			put(duckFinish, y+row, []string{"#.", ".#"}[row%2], "")
		}
		for col := 12 + (18-(tick/3+lane*4)%18)%18; col < duckRaceWidth-2; col += 18 {
			if col+2 < duckFinish || col > duckFinish+1 {
				put(col, y+5, "~.", duckRaceWater)
			}
		}
		if lane < 2 {
			put(0, y+6, strings.Repeat("-", duckRaceWidth), "")
		}

		x := duckRacePosition(tick, lane)
		beat := (tick/3 + lane) % 4
		crest := []string{"          .-==-.", `            \|/`, "            _@"}[lane]
		head := "           ( o)=>"
		wing := []string{`  \  (___)  /`, `  \  (._)   /`, `  \  (\_)   /`, `  \  (._)   /`}[beat]
		body := `   '-------'`
		if lane == 2 && tick >= 51 && tick < 67 {
			head = "           ( -)=>"
			wing = `  \  (___)  /`
		}
		put(x, y, crest, laneColor)
		put(x, y+1, head[:15], laneColor)
		put(x+15, y+1, head[15:], duckRaceBill)
		put(x, y+2, `,         _/  )`, laneColor)
		put(x, y+3, `\'-------'   /`, laneColor)
		put(x, y+4, wing, laneColor)
		put(x, y+5, body, laneColor)
		if tick >= 18 && tick < 105 {
			if lane == 2 && tick >= 67 && tick < 94 {
				put(x-8, y+3, []string{"* .-~  ", "' * ~- "}[beat%2], duckRaceWater)
				put(x-9, y+4, "~~--~~~", duckRaceWater)
			} else if x > 14 {
				put(x-4, y+4, []string{"~--", " ~~", "-~-", " ~~"}[beat], duckRaceWater)
			}
		}
		if lane == 2 {
			switch {
			case tick < 51:
				put(40, y+1, ".--.", duckRaceBill)
				put(40, y+2, "|::|", duckRaceBill)
				put(40, y+3, "'--'", duckRaceBill)
			case tick < 58:
				put(40, y+1, ",.:", duckRaceBill)
				put(40, y+2, " '.", duckRaceBill)
			case tick < 66:
				put(41, y+1, "nom!", duckRaceBill)
			}
		}
		if lane == 2 && tick >= 95 {
			put(x-4, y+1, []string{"*  .", ".  +", "+  .", ".  *"}[beat], duckRaceGold)
		}
		if tick >= 99 {
			put(2, y+4, []string{"2nd", "3rd", "1st"}[lane], laneColor)
		}
	}
	put(0, 24, strings.Repeat("=", duckRaceWidth), "")
	put(1, 26, "RACE CALL  >  "+duckRaceCall(tick), duckRaceGold)
	var frame strings.Builder
	for y := range pond {
		if y > 0 {
			frame.WriteByte('\n')
		}
		line := strings.TrimRight(string(pond[y][:]), " ")
		active := ""
		for x := 0; x < len(line); x++ {
			if color := colors[y][x]; line[x] != ' ' && color != active {
				if color == "" {
					frame.WriteString(duckRaceReset)
				} else {
					frame.WriteString(color)
				}
				active = color
			}
			frame.WriteByte(line[x])
		}
		if active != "" {
			frame.WriteString(duckRaceReset)
		}
	}
	return frame.String()
}
