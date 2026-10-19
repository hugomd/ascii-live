package frames

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	rtWidth  = 80
	rtHeight = 24
	rtFrames = 120
)

var RainbowTunnel = FrameType{
	GetFrame:  getRainbowTunnelFrame,
	GetLength: func() int { return rtFrames },
	GetSleep:  func() time.Duration { return 50 * time.Millisecond },
}

var rtShades = []byte(" .:-=+*#%@")

func getRainbowTunnelFrame(i int) string {
	t := float64(rtWrap(i, rtFrames)) / float64(rtFrames)
	cx := float64(rtWidth-1) / 2
	cy := float64(rtHeight-1) / 2

	var b strings.Builder
	b.Grow(rtWidth * rtHeight * 12)
	last := -1

	for y := 0; y < rtHeight; y++ {
		for x := 0; x < rtWidth; x++ {
			nx := (float64(x) - cx) / cx
			ny := (float64(y) - cy) / cy * 2
			dist := math.Hypot(nx, ny)
			if dist < 0.04 {
				rtWriteCell(&b, &last, '@', rtHSV(math.Mod(t*2, 1), 0.2, 1))
				continue
			}

			depth := 1 / dist
			ang := math.Atan2(ny, nx)
			u := ang / (2 * math.Pi)
			v := depth*0.35 + t*6

			ring := 0.5 + 0.5*math.Sin(v*2*math.Pi)
			spoke := 0.5 + 0.5*math.Cos(u*12*math.Pi-depth*1.4+t*4*math.Pi)
			wall := rtClamp(0.35 + 0.65*ring*ring + 0.25*spoke)

			hue := math.Mod(u+t*0.85+depth*0.04, 1)
			if hue < 0 {
				hue += 1
			}
			near := rtClamp(1.25 - dist*0.7)
			val := rtClamp(0.38 + 0.62*wall*near)
			if ny > 0.15 {
				val *= 0.82
			}

			shade := int(wall * float64(len(rtShades)-1))
			if shade < 1 {
				shade = 1
			} else if shade >= len(rtShades) {
				shade = len(rtShades) - 1
			}
			rtWriteCell(&b, &last, rtShades[shade], rtHSV(hue, 0.95, val))
		}
		if last != -1 {
			b.WriteString("\033[0m")
			last = -1
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func rtWriteCell(b *strings.Builder, last *int, ch byte, color uint8) {
	c := int(color)
	if c != *last {
		b.WriteString("\033[38;5;")
		b.WriteString(strconv.Itoa(c))
		b.WriteByte('m')
		*last = c
	}
	b.WriteByte(ch)
}

func rtHSV(h, s, v float64) uint8 {
	h = math.Mod(h, 1)
	if h < 0 {
		h += 1
	}
	s = rtClamp(s)
	v = rtClamp(v)
	if s <= 0.02 {
		g := int(v*23 + 0.5)
		if g > 23 {
			g = 23
		}
		return uint8(232 + g)
	}
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
	return uint8(16 + 36*rtLevel(r) + 6*rtLevel(g) + rtLevel(b))
}

func rtLevel(c float64) int {
	n := int(c*5 + 0.5)
	if n < 0 {
		return 0
	}
	if n > 5 {
		return 5
	}
	return n
}

func rtWrap(i, n int) int {
	if n <= 0 {
		return 0
	}
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

func rtClamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
