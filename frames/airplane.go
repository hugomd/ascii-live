package frames

import (
	"math"
	"strings"
	"time"
)

const (
	apWidth  = 72
	apHeight = 20
	apFrames = 72
)

// Airplane flies in from the left, pitches upright, completes a full loop,
// then levels out and exits straight to the right.
var Airplane = FrameType{
	GetFrame:  getAirplaneFrame,
	GetLength: func() int { return apFrames },
	GetSleep:  func() time.Duration { return 80 * time.Millisecond },
}

// Triangular plane geometry (same silhouette at every heading).
// apAspect compensates for tall terminal cells so the triangle looks even.
const (
	apNoseLen = 5.0 // tip ahead of center
	apWingLen = 3.5 // base corners behind/side
	apAspect  = 0.55
)

func apWrap(i, n int) int {
	if n <= 0 {
		return 0
	}
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

// apDir maps a heading unit vector into screen cells (y down, aspect-corrected).
func apDir(headingDeg float64) (dx, dy float64) {
	rad := headingDeg * math.Pi / 180
	return math.Cos(rad), -math.Sin(rad) * apAspect
}

// apTriangleTips returns nose and two wing tips for an isosceles triangle
// pointing along heading (0° = right, 90° = up). Same geometry every frame.
func apTriangleTips(cx, cy, headingDeg float64) (nose, left, right [2]float64) {
	hx, hy := apDir(headingDeg)
	px, py := apDir(headingDeg + 90) // wing axis (perpendicular)

	nose = [2]float64{cx + apNoseLen*hx, cy + apNoseLen*hy}
	// Base sits behind the center so the shape stays a forward-pointing triangle.
	bx := cx - (apNoseLen*0.55)*hx
	by := cy - (apNoseLen*0.55)*hy
	left = [2]float64{bx + apWingLen*px, by + apWingLen*py}
	right = [2]float64{bx - apWingLen*px, by - apWingLen*py}
	return nose, left, right
}

func apPlot(canvas [][]rune, x, y int, ch rune) {
	if x >= 0 && x < apWidth && y >= 0 && y < apHeight {
		canvas[y][x] = ch
	}
}

// apDrawLine draws a line with a stable glyph so the triangle edges match
// before, during, and after the maneuver.
func apDrawLine(canvas [][]rune, x0, y0, x1, y1 float64, ch rune) {
	dx := x1 - x0
	dy := y1 - y0
	steps := int(math.Ceil(math.Hypot(dx, dy)))
	if steps < 1 {
		apPlot(canvas, int(math.Round(x0)), int(math.Round(y0)), ch)
		return
	}
	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		apPlot(canvas, int(math.Round(x0+dx*t)), int(math.Round(y0+dy*t)), ch)
	}
}

func apFillTriangle(canvas [][]rune, a, b, c [2]float64) {
	minX := int(math.Floor(math.Min(a[0], math.Min(b[0], c[0]))))
	maxX := int(math.Ceil(math.Max(a[0], math.Max(b[0], c[0]))))
	minY := int(math.Floor(math.Min(a[1], math.Min(b[1], c[1]))))
	maxY := int(math.Ceil(math.Max(a[1], math.Max(b[1], c[1]))))

	area := func(p, q, r [2]float64) float64 {
		return (q[0]-p[0])*(r[1]-p[1]) - (q[1]-p[1])*(r[0]-p[0])
	}
	abc := area(a, b, c)
	if math.Abs(abc) < 1e-6 {
		return
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := [2]float64{float64(x) + 0.5, float64(y) + 0.5}
			w0 := area(b, c, p)
			w1 := area(c, a, p)
			w2 := area(a, b, p)
			if (w0 >= 0 && w1 >= 0 && w2 >= 0 && abc > 0) ||
				(w0 <= 0 && w1 <= 0 && w2 <= 0 && abc < 0) {
				apPlot(canvas, x, y, '#')
			}
		}
	}
}

// Exact left→right cruise silhouette (matches the reference triangle).
// Wide middle tapering to a tip on the right — used only in level flight.
var apLevelSprite = []string{
	`####`,
	`#########`,
	`###########`,
	`####`,
}

func apBlitSprite(canvas [][]rune, sprite []string, cx, cy int) {
	if len(sprite) == 0 {
		return
	}
	maxW := 0
	for _, line := range sprite {
		if len(line) > maxW {
			maxW = len(line)
		}
	}
	ox := cx - maxW/2
	oy := cy - len(sprite)/2
	for row, line := range sprite {
		for col, ch := range line {
			if ch == ' ' {
				continue
			}
			apPlot(canvas, ox+col, oy+row, ch)
		}
	}
}

func apDrawAirplane(canvas [][]rune, cx, cy, headingDeg float64) {
	// Level flight (enter / exit): use the fixed triangular sprite.
	h := headingDeg
	for h < 0 {
		h += 360
	}
	for h >= 360 {
		h -= 360
	}
	if h < 20 || h > 340 {
		apBlitSprite(canvas, apLevelSprite, int(math.Round(cx)), int(math.Round(cy)))
		return
	}

	nose, left, right := apTriangleTips(cx, cy, headingDeg)
	apFillTriangle(canvas, nose, left, right)
	apDrawLine(canvas, nose[0], nose[1], left[0], left[1], '#')
	apDrawLine(canvas, nose[0], nose[1], right[0], right[1], '#')
	apDrawLine(canvas, left[0], left[1], right[0], right[1], '#')
}

// apPose returns canvas position (cx, cy) and heading degrees for progress t in [0,1].
// Path: enter from left → pitch upright into a full loop → exit straight right.
func apPose(t float64) (cx, cy, heading float64) {
	const (
		loopCX   = 34.0
		loopCY   = 9.0
		loopR    = 6.5
		enterEnd = 0.25
		loopEnd  = 0.78
	)
	cruiseY := loopCY + loopR // bottom of the loop

	switch {
	case t < enterEnd:
		// Fly in from the left, level, toward the loop entry.
		// Start far enough in that the full triangular sprite is on-screen.
		p := t / enterEnd
		cx = 8 + p*(loopCX-8)
		cy = cruiseY
		heading = 0
	case t < loopEnd:
		// Full round turn: bottom → upright (right side) → inverted top →
		// down the left → back to bottom, heading right again.
		p := (t - enterEnd) / (loopEnd - enterEnd)
		a := p * 2 * math.Pi
		cx = loopCX + loopR*math.Sin(a)
		cy = loopCY + loopR*math.Cos(a)
		// Tangent heading: 0° at bottom (right), 90° upright, 180° at top, etc.
		heading = a * 180 / math.Pi
	default:
		// Level out and exit straight to the right.
		p := (t - loopEnd) / (1 - loopEnd)
		cx = loopCX + p*(apWidth+6-loopCX)
		cy = cruiseY
		heading = 0
	}
	return cx, cy, heading
}

func getAirplaneFrame(i int) string {
	t := float64(apWrap(i, apFrames)) / float64(apFrames)
	cx, cy, heading := apPose(t)

	canvas := make([][]rune, apHeight)
	for y := 0; y < apHeight; y++ {
		row := make([]rune, apWidth)
		for x := range row {
			row[x] = ' '
		}
		// Soft ground / horizon hint on the last row.
		if y == apHeight-1 {
			for x := range row {
				row[x] = '.'
			}
		}
		canvas[y] = row
	}

	// Contrail dots behind the plane during cruise/exit (level flight).
	if t < 0.25 || t > 0.78 {
		trailLen := 8
		for n := 1; n <= trailLen; n++ {
			tx := int(cx) - n*2
			ty := int(cy)
			if tx >= 0 && tx < apWidth && ty >= 0 && ty < apHeight-1 {
				if canvas[ty][tx] == ' ' {
					canvas[ty][tx] = '.'
				}
			}
		}
	}

	apDrawAirplane(canvas, cx, cy, heading)

	var b strings.Builder
	b.Grow(apWidth * apHeight)
	for y := 0; y < apHeight; y++ {
		b.WriteString(strings.TrimRight(string(canvas[y]), " "))
		b.WriteByte('\n')
	}
	return b.String()
}
