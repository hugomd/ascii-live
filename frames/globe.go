package frames

import (
	"math"
	"strings"
	"time"
)

const (
	glWidth  = 64
	glHeight = 28
	glFrames = 90
	glRadius = 11.5
)

// Globe is a transparent 3D wireframe Earth that rotates about its center.
var Globe = FrameType{
	GetFrame:  getGlobeFrame,
	GetLength: func() int { return glFrames },
	GetSleep:  func() time.Duration { return 70 * time.Millisecond },
}

// Coarse equirectangular land mask (rows = lat N→S, cols = lon -180→180).
// '#' = land, '.' = ocean. Sized for a glass-globe silhouette, not cartography.
var glLand = []string{
	"................................................................",
	"..............####..............................................",
	"..........##########.....................###....................",
	"........##############.................#####....................",
	".......################..............########...####............",
	"......########..########............################............",
	"......######......######...........##################...........",
	".....######........#####..........########...#########..........",
	".....#####..........####.........#######.......#######..........",
	".....####...........####.........######.........######..........",
	"......###...........###..........#####...........#####..........",
	".......##..........###...........####.............#####.........",
	"........##........###.............###..............####.........",
	".........###.....###...............##..............####.........",
	"..........########..................#..............###..........",
	"...........######.................................####..........",
	"............####.................................####...........",
	".............###................................#####...........",
	"..............##...............................######...........",
	"...............#..............................#######...........",
	".............................................#######............",
	"............................................######..............",
	"...........................................#####................",
	"..........................................####..................",
	"................................................................",
	"................................................................",
}

var glShades = []rune(" ·∙•○°oO#@")

func glWrap(i, n int) int {
	if n <= 0 {
		return 0
	}
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

func glRotY(x, y, z, a float64) (float64, float64, float64) {
	c, s := math.Cos(a), math.Sin(a)
	return x*c + z*s, y, -x*s + z*c
}

func glLandAt(lat, lon float64) bool {
	// lat ∈ [-π/2, π/2], lon ∈ [-π, π]
	rows := len(glLand)
	if rows == 0 {
		return false
	}
	cols := len(glLand[0])
	v := (0.5 - lat/math.Pi) * float64(rows) // N→S
	u := (lon/math.Pi + 1) * 0.5 * float64(cols)
	ri := int(math.Floor(v))
	ci := int(math.Floor(u))
	if ri < 0 || ri >= rows {
		return false
	}
	row := glLand[ri]
	if len(row) == 0 {
		return false
	}
	ci = ((ci % len(row)) + len(row)) % len(row)
	return row[ci] == '#'
}

func glNearStep(v, step, tol float64) bool {
	m := math.Mod(math.Abs(v), step)
	return m < tol || m > step-tol
}

func glGridAt(lat, lon float64) (onLat, onLon bool) {
	// 45° graticule — sparse enough to read as wireframe in ASCII.
	const step = math.Pi / 4
	onLat = glNearStep(lat, step, 0.022)
	onLon = glNearStep(lon, step, 0.022)
	return onLat, onLon
}

func glShade(n float64) rune {
	if n < 0 {
		n = 0
	}
	if n > 1 {
		n = 1
	}
	i := int(n * float64(len(glShades)-1))
	return glShades[i]
}

func glPlot(buf [][]rune, depth [][]float64, x, y int, z float64, ch rune, force bool) {
	if x < 0 || x >= glWidth || y < 0 || y >= glHeight {
		return
	}
	if force || z < depth[y][x] {
		depth[y][x] = z
		buf[y][x] = ch
	}
}

func glPlotGlass(buf [][]rune, depth [][]float64, x, y int, z float64, ch rune) {
	if x < 0 || x >= glWidth || y < 0 || y >= glHeight {
		return
	}
	// Prefer the nearer (front) face; fill empty cells with the back face for glass.
	if buf[y][x] == ' ' || z < depth[y][x] {
		depth[y][x] = z
		buf[y][x] = ch
	}
}

func getGlobeFrame(i int) string {
	yaw := float64(glWrap(i, glFrames)) * (2 * math.Pi / float64(glFrames))

	buf := make([][]rune, glHeight)
	depth := make([][]float64, glHeight)
	for y := 0; y < glHeight; y++ {
		buf[y] = make([]rune, glWidth)
		depth[y] = make([]float64, glWidth)
		for x := 0; x < glWidth; x++ {
			buf[y][x] = ' '
			depth[y][x] = 1e9
		}
	}

	cx := float64(glWidth) / 2
	cy := float64(glHeight)/2 - 1.2 // slightly above center (mean position)
	aspect := 2.15                  // terminal cell aspect

	// Sample the sphere: front and back hemispheres for glass transparency.
	for sy := 0; sy < glHeight-3; sy++ {
		for sx := 0; sx < glWidth; sx++ {
			nx := (float64(sx) - cx) / (glRadius * aspect)
			ny := (float64(sy) - cy) / glRadius
			rr := nx*nx + ny*ny
			if rr > 1 {
				continue
			}
			nz := math.Sqrt(1 - rr)

			for _, sign := range []float64{-1, 1} { // back first, then front (glass)
				x, y, z := nx, ny, nz*sign
				// Rotate about vertical axis (mean/center position).
				x, y, z = glRotY(x, y, z, yaw)

				lat := math.Asin(math.Max(-1, math.Min(1, y)))
				lon := math.Atan2(x, z)

				// +nz faces the camera; nearer front face wins over the back.
				viewZ := -sign * nz

				onLat, onLon := glGridAt(lat, lon)
				onGrid := onLat || onLon
				onLand := glLandAt(lat, lon)
				var ch rune
				switch {
				case onGrid:
					// Meridians/parallels: lighter on the back for glass depth.
					bright := 0.75
					if sign < 0 {
						bright = 0.28
					}
					ch = glShade(bright)
					if onLat && onLon {
						ch = '+' // graticule crossings
					}
				case onLand:
					bright := 0.5
					if sign < 0 {
						bright = 0.16
					}
					ch = glShade(bright)
				default:
					// Mostly empty glass — limb glow only.
					edge := 1 - rr
					if edge < 0.07 {
						bright := (0.07 - edge) * 5
						if sign < 0 {
							bright *= 0.4
						}
						ch = glShade(bright)
					} else if sign > 0 && edge < 0.15 {
						ch = '·'
					} else {
						continue
					}
				}

				// Interior sparkles (data nodes) on the near hemisphere.
				if sign > 0 && (onLand || onGrid) {
					h := math.Sin(lon*11+yaw*3) * math.Cos(lat*8+yaw*1.5)
					if h > 0.93 {
						ch = '*'
					}
				}

				glPlotGlass(buf, depth, sx, sy, viewZ, ch)
			}
		}
	}

	// Pole highlights (north / south glow).
	for _, pole := range []struct{ y, boost float64 }{{cy - glRadius, 1}, {cy + glRadius, 1}} {
		px := int(math.Round(cx))
		py := int(math.Round(pole.y))
		for dy := -1; dy <= 1; dy++ {
			for dx := -2; dx <= 2; dx++ {
				glPlot(buf, depth, px+dx, py+dy, -2, '*', true)
			}
		}
		glPlot(buf, depth, px, py, -3, '@', true)
	}

	// Glowing circular base / pedestal under the globe.
	baseY := int(math.Round(cy + glRadius + 1.2))
	for dx := -10; dx <= 10; dx++ {
		x := int(cx) + dx
		t := 1 - math.Abs(float64(dx))/10
		ch := glShade(0.25 + 0.55*t)
		if t > 0.85 {
			ch = '@'
		}
		glPlot(buf, depth, x, baseY, 0, ch, true)
		if dx >= -8 && dx <= 8 {
			glPlot(buf, depth, x, baseY+1, 1, glShade(0.15+0.3*t), true)
		}
	}
	// Soft ground shadow ellipse.
	shadowY := baseY + 2
	for dx := -14; dx <= 14; dx++ {
		x := int(cx) + dx
		if x < 0 || x >= glWidth || shadowY >= glHeight {
			continue
		}
		if buf[shadowY][x] == ' ' {
			buf[shadowY][x] = '·'
		}
	}

	var b strings.Builder
	b.Grow(glWidth * glHeight)
	for y := 0; y < glHeight; y++ {
		b.WriteString(strings.TrimRight(string(buf[y]), " "))
		b.WriteByte('\n')
	}
	return b.String()
}
