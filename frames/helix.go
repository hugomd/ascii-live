package frames

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	hxWidth  = 72
	hxHeight = 28
	hxFrames = 90
	hxSteps  = 96
)

var Helix = FrameType{
	GetFrame:  getHelixFrame,
	GetLength: func() int { return hxFrames },
	GetSleep:  func() time.Duration { return 55 * time.Millisecond },
}

type hxCell struct {
	ch    byte
	color uint8
	z     float64
}

func getHelixFrame(i int) string {
	t := float64(hxWrap(i, hxFrames)) * (2 * math.Pi / float64(hxFrames))
	ax := 0.35 + 0.12*math.Sin(t*0.6)
	ay := t
	az := 0.12 * math.Sin(t*0.45)

	grid := make([][]hxCell, hxHeight)
	for y := range grid {
		grid[y] = make([]hxCell, hxWidth)
		for x := range grid[y] {
			grid[y][x].z = 1e9
			grid[y][x].ch = ' '
		}
	}

	strandA := make([]hxPT, hxSteps)
	strandB := make([]hxPT, hxSteps)

	for s := 0; s < hxSteps; s++ {
		u := float64(s) / float64(hxSteps-1)
		theta := u*4*math.Pi + t*0.35
		z := (u - 0.5) * 2.7
		r := 1.05
		a := hxVec{r * math.Cos(theta), r * math.Sin(theta), z}
		b := hxVec{r * math.Cos(theta+math.Pi), r * math.Sin(theta+math.Pi), z}
		strandA[s] = hxProject(hxRotate(a, ax, ay, az))
		strandB[s] = hxProject(hxRotate(b, ax, ay, az))
	}

	for s := 0; s < hxSteps; s++ {
		pa, pb := strandA[s], strandB[s]
		if pa.ok {
			depth := hxShade(pa.z)
			hxDot(grid, pa.x, pa.y, pa.z, hxChar(depth, true), hxHSV(0.55, 0.95, depth))
		}
		if pb.ok {
			depth := hxShade(pb.z)
			hxDot(grid, pb.x, pb.y, pb.z, hxChar(depth, true), hxHSV(0.88, 0.95, depth))
		}
		if s%6 == 0 && pa.ok && pb.ok {
			hue := math.Mod(0.12+float64(s)*0.03+t*0.1, 1)
			hxLine(grid, pa, pb, '=', hxHSV(hue, 0.9, 0.85))
		}
		if s+1 < hxSteps {
			if strandA[s].ok && strandA[s+1].ok {
				hxLine(grid, strandA[s], strandA[s+1], '@', hxHSV(0.55, 0.95, hxShade(strandA[s].z)))
			}
			if strandB[s].ok && strandB[s+1].ok {
				hxLine(grid, strandB[s], strandB[s+1], '@', hxHSV(0.88, 0.95, hxShade(strandB[s].z)))
			}
		}
	}

	var b strings.Builder
	b.Grow(hxWidth * hxHeight * 10)
	last := -1
	for y := 0; y < hxHeight; y++ {
		for x := 0; x < hxWidth; x++ {
			c := grid[y][x]
			if c.ch == ' ' {
				if last != -1 {
					b.WriteString("\033[0m")
					last = -1
				}
				b.WriteByte(' ')
				continue
			}
			col := int(c.color)
			if col != last {
				b.WriteString("\033[38;5;")
				b.WriteString(strconv.Itoa(col))
				b.WriteByte('m')
				last = col
			}
			b.WriteByte(c.ch)
		}
		if last != -1 {
			b.WriteString("\033[0m")
			last = -1
		}
		b.WriteByte('\n')
	}
	return b.String()
}

type hxVec struct{ x, y, z float64 }

func hxRotate(v hxVec, ax, ay, az float64) hxVec {
	cx, sx := math.Cos(ax), math.Sin(ax)
	y, z := v.y*cx-v.z*sx, v.y*sx+v.z*cx
	v.y, v.z = y, z
	cy, sy := math.Cos(ay), math.Sin(ay)
	x := v.x*cy + v.z*sy
	z = -v.x*sy + v.z*cy
	v.x, v.z = x, z
	cz, sz := math.Cos(az), math.Sin(az)
	x, y = v.x*cz-v.y*sz, v.x*sz+v.y*cz
	v.x, v.y = x, y
	return v
}

type hxPT = struct {
	x, y, z float64
	ok      bool
}

func hxProject(v hxVec) hxPT {
	v.z += 4.4
	if v.z < 1.2 {
		return hxPT{}
	}
	scale := 34.0
	sx := float64(hxWidth)/2 + (v.x/v.z)*scale*2.15
	sy := float64(hxHeight)/2 + (v.y/v.z)*scale
	if sx < -1 || sy < -1 || sx > float64(hxWidth)+1 || sy > float64(hxHeight)+1 {
		return hxPT{}
	}
	return hxPT{sx, sy, v.z, true}
}

func hxDot(grid [][]hxCell, x, y, z float64, ch byte, color uint8) {
	hxPlot(grid, x, y, z, ch, color)
	hxPlot(grid, x+1, y, z+0.001, ch, color)
}

func hxPlot(grid [][]hxCell, x, y, z float64, ch byte, color uint8) {
	xi, yi := int(math.Round(x)), int(math.Round(y))
	if xi < 0 || yi < 0 || xi >= hxWidth || yi >= hxHeight {
		return
	}
	if z <= grid[yi][xi].z {
		grid[yi][xi] = hxCell{ch, color, z}
	}
}

func hxLine(grid [][]hxCell, a, b hxPT, ch byte, color uint8) {
	dx := b.x - a.x
	dy := b.y - a.y
	if math.Hypot(dx, dy) > 24 {
		return
	}
	steps := int(math.Max(math.Abs(dx), math.Abs(dy)))
	if steps < 1 {
		hxDot(grid, a.x, a.y, a.z, ch, color)
		return
	}
	for i := 0; i <= steps; i++ {
		u := float64(i) / float64(steps)
		hxDot(grid, a.x+dx*u, a.y+dy*u, a.z+(b.z-a.z)*u, ch, color)
	}
}

func hxShade(z float64) float64 {
	return hxClamp(1.15 - (z-3.2)*0.35)
}

func hxChar(depth float64, backbone bool) byte {
	if backbone {
		if depth > 0.8 {
			return '@'
		}
		if depth > 0.55 {
			return '#'
		}
		return '*'
	}
	return '='
}

func hxHSV(h, s, v float64) uint8 {
	h = math.Mod(h, 1)
	if h < 0 {
		h += 1
	}
	s = hxClamp(s)
	v = hxClamp(v)
	var r, g, b float64
	hp := h * 6
	x := 1 - math.Abs(math.Mod(hp, 2)-1)
	switch {
	case hp < 1:
		r, g, b = 1, x, 0
	case hp < 2:
		r, g, b = x, 1, 0
	case hp < 3:
		r, g, b = 0, 1, x
	case hp < 4:
		r, g, b = 0, x, 1
	case hp < 5:
		r, g, b = x, 0, 1
	default:
		r, g, b = 1, 0, x
	}
	r = (r*s + 1 - s) * v
	g = (g*s + 1 - s) * v
	b = (b*s + 1 - s) * v
	return uint8(16 + 36*hxLevel(r) + 6*hxLevel(g) + hxLevel(b))
}

func hxLevel(c float64) int {
	n := int(c*5 + 0.5)
	if n < 0 {
		return 0
	}
	if n > 5 {
		return 5
	}
	return n
}

func hxWrap(i, n int) int {
	if n <= 0 {
		return 0
	}
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

func hxClamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
