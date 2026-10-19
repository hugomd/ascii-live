package frames

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	auWidth  = 80
	auHeight = 24
	auFrames = 100
)

var Aurora = FrameType{
	GetFrame:  getAuroraFrame,
	GetLength: func() int { return auFrames },
	GetSleep:  func() time.Duration { return 65 * time.Millisecond },
}

var auShades = []byte(" .:-=+*#%@")

func getAuroraFrame(i int) string {
	t := float64(auWrap(i, auFrames)) * (2 * math.Pi / float64(auFrames))

	var b strings.Builder
	b.Grow(auWidth * auHeight * 12)
	last := -1

	for y := 0; y < auHeight; y++ {
		ny := float64(y) / float64(auHeight-1)
		for x := 0; x < auWidth; x++ {
			nx := float64(x) / float64(auWidth-1)

			best := 0.0
			hue := 0.38
			depth := 0.0

			for layer := 0; layer < 4; layer++ {
				lf := float64(layer)
				phase := t*(0.55+lf*0.18) + lf*1.3
				fold := math.Sin(nx*(2.4+lf*0.7)+phase) + 0.35*math.Sin(nx*(5.1-lf)+phase*1.6)
				ridge := 0.10 + 0.13*lf + 0.10*fold
				dist := ny - ridge
				if dist < -0.08 {
					continue
				}
				curtain := math.Exp(-(dist*dist)/(0.018+0.03*lf)) * (1.05 - ny*0.72)
				ripple := 0.55 + 0.45*math.Sin(nx*9+ny*6+phase+lf)
				near := 1 - lf*0.16
				glow := curtain * ripple * near
				if glow > best {
					best = glow
					depth = near
					hue = math.Mod(0.32+lf*0.14+0.06*fold+t*0.04, 1)
				}
			}

			star := auStar(x, y, t)
			if best < 0.08 {
				if star > 0 {
					auWriteCell(&b, &last, star, auHSV(0.62, 0.15, 0.85))
				} else {
					auWriteCell(&b, &last, ' ', 16)
				}
				continue
			}

			val := auClamp(0.28 + 0.72*best*depth)
			sat := 0.75 + 0.2*best
			shade := int(best * float64(len(auShades)-1))
			if shade < 1 {
				shade = 1
			} else if shade >= len(auShades) {
				shade = len(auShades) - 1
			}
			ch := auShades[shade]
			if ny > 0.72 && best < 0.35 {
				ch = '.'
			}
			if star == '*' && best < 0.22 {
				ch = '*'
			}
			auWriteCell(&b, &last, ch, auHSV(hue, sat, val))
		}
		if last != -1 {
			b.WriteString("\033[0m")
			last = -1
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func auStar(x, y int, t float64) byte {
	n := (x*47 + y*113) & 255
	if n > 6 {
		return 0
	}
	twinkle := 0.5 + 0.5*math.Sin(t*3+float64(n)*1.7)
	if twinkle < 0.35 {
		return 0
	}
	if n <= 2 && twinkle > 0.8 {
		return '*'
	}
	return '.'
}

func auWriteCell(b *strings.Builder, last *int, ch byte, color uint8) {
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

func auHSV(h, s, v float64) uint8 {
	h = math.Mod(h, 1)
	if h < 0 {
		h += 1
	}
	s = auClamp(s)
	v = auClamp(v)
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
	return uint8(16 + 36*auLevel(r) + 6*auLevel(g) + auLevel(b))
}

func auLevel(c float64) int {
	n := int(c*5 + 0.5)
	if n < 0 {
		return 0
	}
	if n > 5 {
		return 5
	}
	return n
}

func auWrap(i, n int) int {
	if n <= 0 {
		return 0
	}
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

func auClamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
