package frames

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	cbWidth  = 64
	cbHeight = 26
	cbFrames = 90
)

var CubeLite = FrameType{
	GetFrame:  getCubeLiteFrame,
	GetLength: func() int { return cbFrames },
	GetSleep:  func() time.Duration { return 55 * time.Millisecond },
}

var cbShades = []byte(" .:-=+*#%@")

var cbFaceHue = [6]float64{0.00, 0.16, 0.33, 0.50, 0.66, 0.83}

var cbFaces = [6][4]int{
	{4, 5, 6, 7},
	{1, 0, 3, 2},
	{5, 1, 2, 6},
	{0, 4, 7, 3},
	{7, 6, 2, 3},
	{0, 1, 5, 4},
}

type cbVec struct{ x, y, z float64 }

func getCubeLiteFrame(i int) string {
	t := float64(cbWrap(i, cbFrames)) * (2 * math.Pi / float64(cbFrames))
	ax := 0.62
	ay := math.Pi/4 + t
	az := 0.08 * math.Sin(t*0.8)

	verts := [8]cbVec{
		{-1, -1, -1}, {1, -1, -1}, {1, 1, -1}, {-1, 1, -1},
		{-1, -1, 1}, {1, -1, 1}, {1, 1, 1}, {-1, 1, 1},
	}
	var world [8]cbVec
	var proj [8]cbVec
	for n, v := range verts {
		w := cbRotate(v, ax, ay, az)
		world[n] = w
		w.z += 4.6
		proj[n] = cbVec{
			float64(cbWidth)/2 + (w.x/w.z)*24*2.15,
			float64(cbHeight)/2 + (w.y/w.z)*24 - 0.8,
			w.z,
		}
	}

	grid := make([][]cellT, cbHeight)
	for y := range grid {
		grid[y] = make([]cellT, cbWidth)
		for x := range grid[y] {
			grid[y][x].z = 1e9
			grid[y][x].ch = ' '
		}
	}

	light := cbNorm(cbVec{0.55, -0.75, -0.4})
	cam := cbVec{0, 0, -4.6}

	for fi, face := range cbFaces {
		a, b, c, d := world[face[0]], world[face[1]], world[face[2]], world[face[3]]
		n := cbNorm(cbCross(cbSub(b, a), cbSub(d, a)))
		center := cbVec{(a.x + b.x + c.x + d.x) / 4, (a.y + b.y + c.y + d.y) / 4, (a.z + b.z + c.z + d.z) / 4}
		if cbDot(n, cbSub(cam, center)) <= 0.02 {
			continue
		}
		ndot := cbDot(n, light)
		if ndot < 0 {
			ndot = 0
		}
		val := cbClamp(0.45 + 0.55*ndot)
		shade := int(val * float64(len(cbShades)-1))
		if shade < 3 {
			shade = 3
		} else if shade >= len(cbShades) {
			shade = len(cbShades) - 1
		}
		color := cbHSV(cbFaceHue[fi], 0.92, val)
		ch := cbShades[shade]
		cbFill(grid, proj[face[0]], proj[face[1]], proj[face[2]], ch, color)
		cbFill(grid, proj[face[0]], proj[face[2]], proj[face[3]], ch, color)
	}

	edgeColor := cbHSV(0.12, 0.15, 1)
	edges := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {4, 5}, {5, 6}, {6, 7}, {7, 4}, {0, 4}, {1, 5}, {2, 6}, {3, 7}}
	for _, e := range edges {
		cbLine(grid, proj[e[0]], proj[e[1]], '#', edgeColor)
	}

	sumx := 0.0
	for _, p := range proj {
		sumx += p.x
	}
	shadowX := sumx / 8
	shadowY := float64(cbHeight) - 2.3
	for y := cbHeight - 4; y < cbHeight-1; y++ {
		for x := 0; x < cbWidth; x++ {
			dx := (float64(x) - shadowX) / 13
			dy := (float64(y) - shadowY) / 1.3
			if dx*dx+dy*dy >= 1 || grid[y][x].ch != ' ' {
				continue
			}
			fall := 1 - (dx*dx + dy*dy)
			grid[y][x] = cellT{'.', cbHSV(0.62, 0.25, 0.22+0.28*fall), 9}
		}
	}

	var b strings.Builder
	b.Grow(cbWidth * cbHeight * 10)
	last := -1
	for y := 0; y < cbHeight; y++ {
		for x := 0; x < cbWidth; x++ {
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

func cbRotate(v cbVec, ax, ay, az float64) cbVec {
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

func cbSub(a, b cbVec) cbVec { return cbVec{a.x - b.x, a.y - b.y, a.z - b.z} }
func cbDot(a, b cbVec) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}
func cbCross(a, b cbVec) cbVec {
	return cbVec{a.y*b.z - a.z*b.y, a.z*b.x - a.x*b.z, a.x*b.y - a.y*b.x}
}
func cbNorm(v cbVec) cbVec {
	l := math.Sqrt(cbDot(v, v))
	if l == 0 {
		return v
	}
	return cbVec{v.x / l, v.y / l, v.z / l}
}

func cbFill(grid [][]cellT, a, b, c cbVec, ch byte, color uint8) {
	minx := int(math.Floor(math.Min(a.x, math.Min(b.x, c.x))))
	maxx := int(math.Ceil(math.Max(a.x, math.Max(b.x, c.x))))
	miny := int(math.Floor(math.Min(a.y, math.Min(b.y, c.y))))
	maxy := int(math.Ceil(math.Max(a.y, math.Max(b.y, c.y))))
	if minx < 0 {
		minx = 0
	}
	if miny < 0 {
		miny = 0
	}
	if maxx >= cbWidth {
		maxx = cbWidth - 1
	}
	if maxy >= cbHeight {
		maxy = cbHeight - 1
	}
	denom := (b.y-c.y)*(a.x-c.x) + (c.x-b.x)*(a.y-c.y)
	if math.Abs(denom) < 1e-6 {
		return
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			px := float64(x) + 0.5
			py := float64(y) + 0.5
			w0 := ((b.y-c.y)*(px-c.x) + (c.x-b.x)*(py-c.y)) / denom
			w1 := ((c.y-a.y)*(px-c.x) + (a.x-c.x)*(py-c.y)) / denom
			w2 := 1 - w0 - w1
			if w0 < -0.01 || w1 < -0.01 || w2 < -0.01 {
				continue
			}
			z := w0*a.z + w1*b.z + w2*c.z
			if z < grid[y][x].z {
				grid[y][x].ch = ch
				grid[y][x].color = color
				grid[y][x].z = z
			}
		}
	}
}

func cbLine(grid [][]cellT, a, b cbVec, ch byte, color uint8) {
	dx := b.x - a.x
	dy := b.y - a.y
	steps := int(math.Max(math.Abs(dx), math.Abs(dy)))
	if steps < 1 {
		cbPlot(grid, a.x, a.y, a.z-0.02, ch, color)
		return
	}
	for i := 0; i <= steps; i++ {
		u := float64(i) / float64(steps)
		cbPlot(grid, a.x+dx*u, a.y+dy*u, a.z+(b.z-a.z)*u-0.02, ch, color)
	}
}

func cbPlot(grid [][]cellT, x, y, z float64, ch byte, color uint8) {
	xi, yi := int(math.Round(x)), int(math.Round(y))
	if xi < 0 || yi < 0 || xi >= cbWidth || yi >= cbHeight {
		return
	}
	if z <= grid[yi][xi].z {
		grid[yi][xi].ch = ch
		grid[yi][xi].color = color
		grid[yi][xi].z = z
	}
}

// cellT matches the anonymous cell used in the raster buffers.
type cellT = struct {
	ch    byte
	color uint8
	z     float64
}

func cbHSV(h, s, v float64) uint8 {
	h = math.Mod(h, 1)
	if h < 0 {
		h += 1
	}
	s = cbClamp(s)
	v = cbClamp(v)
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
	return uint8(16 + 36*cbLevel(r) + 6*cbLevel(g) + cbLevel(b))
}

func cbLevel(c float64) int {
	n := int(c*5 + 0.5)
	if n < 0 {
		return 0
	}
	if n > 5 {
		return 5
	}
	return n
}

func cbWrap(i, n int) int {
	if n <= 0 {
		return 0
	}
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

func cbClamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
