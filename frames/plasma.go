package frames

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	plWidth  = 80
	plHeight = 24
	plFrames = 90
)

var Plasma = FrameType{
	GetFrame:  getPlasmaFrame,
	GetLength: func() int { return plFrames },
	GetSleep:  func() time.Duration { return 50 * time.Millisecond },
}

var plShades = []byte(" .:-=+*#%@")

func getPlasmaFrame(i int) string {
	t := float64(plWrap(i, plFrames)) * (2 * math.Pi / float64(plFrames))
	cx := float64(plWidth-1) / 2
	cy := float64(plHeight-1) / 2
	rx := cx * 0.72
	ry := cy * 0.92

	var b strings.Builder
	b.Grow(plWidth * plHeight * 12)
	last := -1

	lx, ly, lz := -0.45, -0.35, 0.82
	inv := 1 / math.Sqrt(lx*lx+ly*ly+lz*lz)
	lx, ly, lz = lx*inv, ly*inv, lz*inv

	for y := 0; y < plHeight; y++ {
		for x := 0; x < plWidth; x++ {
			nx := (float64(x) - cx) / rx
			ny := (float64(y) - cy) / ry
			r2 := nx*nx + ny*ny

			if r2 > 1.35 {
				plWriteCell(&b, &last, ' ', 16)
				continue
			}

			var z, halo float64
			if r2 <= 1 {
				z = math.Sqrt(1 - r2)
				halo = 0
			} else {
				z = 0
				halo = 1 - (r2-1)/0.35
			}

			field := math.Sin(nx*3.1 + t)
			field += math.Sin(ny*3.6 - t*1.2)
			field += math.Sin((nx+ny)*2.4 + t*0.8)
			field += math.Sin((z*4.2 + nx*ny*3) + t*1.4)
			field *= 0.25

			hue := math.Mod(0.55+field*0.45+t*0.08, 1)
			if hue < 0 {
				hue += 1
			}

			ndot := nx*lx + ny*ly + z*lz
			if ndot < 0 {
				ndot = 0
			}
			lit := 0.22 + 0.78*ndot
			if halo > 0 {
				lit = 0.18 + 0.55*halo
				hue = math.Mod(hue+0.08, 1)
			}

			val := plClamp(lit * (0.55 + 0.45*(field+1)*0.5))
			sat := 0.85
			if halo > 0 {
				sat = 0.55 + 0.3*halo
				val *= halo
			}

			shade := int((0.25 + 0.75*val) * float64(len(plShades)-1))
			if shade < 1 {
				shade = 1
			} else if shade >= len(plShades) {
				shade = len(plShades) - 1
			}
			ch := plShades[shade]
			if halo > 0 && halo < 0.45 {
				ch = '.'
			}
			plWriteCell(&b, &last, ch, plHSV(hue, sat, plClamp(val)))
		}
		if last != -1 {
			b.WriteString("\033[0m")
			last = -1
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func plWriteCell(b *strings.Builder, last *int, ch byte, color uint8) {
	if ch == ' ' {
		if *last != -1 {
			b.WriteString("\033[0m")
			*last = -1
		}
		b.WriteByte(' ')
		return
	}
	c := int(color)
	if c != *last {
		b.WriteString("\033[38;5;")
		b.WriteString(strconv.Itoa(c))
		b.WriteByte('m')
		*last = c
	}
	b.WriteByte(ch)
}

func plHSV(h, s, v float64) uint8 {
	h = math.Mod(h, 1)
	if h < 0 {
		h += 1
	}
	s = plClamp(s)
	v = plClamp(v)
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
	return uint8(16 + 36*plLevel(r) + 6*plLevel(g) + plLevel(b))
}

func plLevel(c float64) int {
	n := int(c*5 + 0.5)
	if n < 0 {
		return 0
	}
	if n > 5 {
		return 5
	}
	return n
}

func plWrap(i, n int) int {
	if n <= 0 {
		return 0
	}
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

func plClamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
