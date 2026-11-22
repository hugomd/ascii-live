package frames

import (
	"math"
	"strings"
)

var Chip = generateChip()

func generateChip() []string {
	var out []string
	w, h, n := 100, 27, 24
	type v3 struct{ x, y, z float64 }
	a0 := v3{0, -1.25, 0.15}
	b0 := v3{-1.15, 0.9, -0.15}
	c0 := v3{1.15, 0.9, -0.15}
	shades := ".,-~:;=!*#$@"

	rot := func(v v3, cy, sy, cx, sx float64) v3 {
		x1, z1 := v.x*cy+v.z*sy, -v.x*sy+v.z*cy
		return v3{x1, v.y*cx - z1*sx, v.y*sx + z1*cx}
	}

	proj := func(v v3) (float64, float64) {
		s := 18.0 / (v.z + 3.2)
		return float64(w)/2 + v.x*s*2, float64(h)/2 + v.y*s
	}

	for f := 0; f < n; f++ {
		ay := float64(f) * 2 * math.Pi / float64(n)
		ax := 0.3 * math.Sin(ay)
		cy, sy, cx, sx := math.Cos(ay), math.Sin(ay), math.Cos(ax), math.Sin(ax)

		pa, pb, pc := rot(a0, cy, sy, cx, sx), rot(b0, cy, sy, cx, sx), rot(c0, cy, sy, cx, sx)

		ux, uy, uz := pb.x-pa.x, pb.y-pa.y, pb.z-pa.z
		vx, vy, vz := pc.x-pa.x, pc.y-pa.y, pc.z-pa.z
		nx, ny, nz := uy*vz-uz*vy, uz*vx-ux*vz, ux*vy-uy*vx
		l := math.Sqrt(nx*nx + ny*ny + nz*nz)
		lum := math.Abs((nx*-0.4 + ny*-0.5 + nz*0.8) / l)
		if lum > 1 {
			lum = 1
		}
		ch := shades[int(lum*float64(len(shades)-1))]

		xa, ya := proj(pa)
		xb, yb := proj(pb)
		xc, yc := proj(pc)

		var sb strings.Builder
		sb.WriteString("\033[93m")
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				px, py := float64(x), float64(y)
				d1 := (px-xb)*(ya-yb) - (xa-xb)*(py-yb)
				d2 := (px-xc)*(yb-yc) - (xb-xc)*(py-yc)
				d3 := (px-xa)*(yc-ya) - (xc-xa)*(py-ya)
				if !((d1 < 0 || d2 < 0 || d3 < 0) && (d1 > 0 || d2 > 0 || d3 > 0)) {
					sb.WriteByte(ch)
				} else {
					sb.WriteByte(' ')
				}
			}
			sb.WriteByte('\n')
		}
		sb.WriteString("\033[0m")
		out = append(out, sb.String())
	}
	return out
}
