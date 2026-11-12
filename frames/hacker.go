package frames

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	hkWidth  = 72
	hkHeight = 24
	hkFrames = 150
)

// Hacker is a three-scene hacker terminal: boot log, ACCESS GRANTED banner,
// and a skull that ends with SYSTEM OWNED.
var Hacker = FrameType{
	GetFrame:  getHackerFrame,
	GetLength: func() int { return hkFrames },
	GetSleep:  func() time.Duration { return 55 * time.Millisecond },
}

func getHackerFrame(i int) string {
	p := hkWrap(i, hkFrames)
	switch {
	case p < 50:
		return hkBoot(p)
	case p < 100:
		return hkBanner(p - 50)
	default:
		return hkSkull(p - 100)
	}
}

func hkWrap(i, n int) int {
	i %= n
	if i < 0 {
		i += n
	}
	return i
}

func hkCanvas() [][]rune {
	g := make([][]rune, hkHeight)
	for y := range g {
		g[y] = []rune(strings.Repeat(" ", hkWidth))
	}
	return g
}

func hkPut(g [][]rune, x, y int, s string) {
	if y < 0 || y >= len(g) {
		return
	}
	for _, ch := range s {
		if x >= 0 && x < len(g[y]) {
			g[y][x] = ch
		}
		x++
	}
}

func hkCenter(g [][]rune, y int, s string) {
	hkPut(g, (len(g[0])-len([]rune(s)))/2, y, s)
}

func hkBox(g [][]rune) {
	w, h := len(g[0]), len(g)
	hkPut(g, 0, 0, "┌"+strings.Repeat("─", w-2)+"┐")
	for y := 1; y < h-1; y++ {
		g[y][0] = '│'
		g[y][w-1] = '│'
	}
	hkPut(g, 0, h-1, "└"+strings.Repeat("─", w-2)+"┘")
}

func hkGlitch(g [][]rune, seed int64, density int) {
	rng := rand.New(rand.NewSource(seed))
	glyphs := []rune("!@#$%&*?/\\|<>[]{}01")
	for y := range g {
		for x := range g[y] {
			if g[y][x] != ' ' && rng.Intn(100) < density {
				g[y][x] = glyphs[rng.Intn(len(glyphs))]
			}
		}
	}
}

func hkJoin(g [][]rune) string {
	lines := make([]string, len(g))
	for y := range g {
		lines[y] = strings.TrimRight(string(g[y]), " ")
	}
	return strings.Join(lines, "\n")
}

func hkGreen(s string) string { return "\033[38;5;46m" + s + "\033[0m" }
func hkRed(s string) string   { return "\033[38;5;196m" + s + "\033[0m" }

func hkBoot(p int) string {
	g := hkCanvas()
	hkBox(g)
	hkCenter(g, 1, " root@mainframe : ~/exploit ")
	lines := []string{
		"root@mainframe:~$ ./h4ck --target 10.13.37.1",
		"[*] resolving target .......... ok",
		"[*] scanning tcp ports ........",
		"    22/tcp   open    ssh",
		"    80/tcp   open    http",
		"    443/tcp  open    https",
		"[*] building exploit chain .... ok",
		"[*] injecting payload ......... ok",
		"[*] escalating privileges ..... ok",
		"[*] hiding tracks ............. ok",
	}
	visible := 1 + p*len(lines)/50
	if visible > len(lines) {
		visible = len(lines)
	}
	for i := 0; i < visible; i++ {
		hkPut(g, 2, 3+i, lines[i])
	}
	cursor := "_"
	if p%2 == 0 {
		cursor = "█"
	}
	hkPut(g, 2+len(lines[visible-1]), 3+visible-1, cursor)

	fill := p * 24 / 50
	pct := p * 2
	if pct > 100 {
		pct = 100
	}
	bar := "[" + strings.Repeat("█", fill) + strings.Repeat("░", 24-fill) + "]"
	hkPut(g, 2, hkHeight-3, bar+fmt.Sprintf(" %3d%%", pct))
	hkPut(g, 2, hkHeight-2, "exploit@10.13.37.1  kernel 5.15.0-x  uid=0(root)")

	return hkGreen(hkJoin(g))
}

func hkBanner(p int) string {
	g := hkCanvas()
	hkBox(g)
	hkCenter(g, 2, "[ ACCESS GRANTED ]")
	y := 4
	for _, line := range strings.Split(hkAccessArt, "\n") {
		hkCenter(g, y, line)
		y++
	}
	y++
	for _, line := range strings.Split(hkGrantedArt, "\n") {
		hkCenter(g, y, line)
		y++
	}
	hkCenter(g, hkHeight-3, "[ system compromised ]")
	if p%24 < 3 {
		hkGlitch(g, int64(p)+7, 6)
	}
	out := hkJoin(g)
	if p%24 == 0 {
		return hkRed(out)
	}
	return hkGreen(out)
}

func hkSkull(p int) string {
	g := hkCanvas()
	hkBox(g)
	hkCenter(g, 1, "[ TARGET ACQUIRED ]")
	for i, line := range strings.Split(hkSkullArt, "\n") {
		hkPut(g, 5, 5+i, line)
	}
	for i, line := range strings.Split(hkSystemArt, "\n") {
		hkPut(g, 33, 6+i, line)
	}
	for i, line := range strings.Split(hkOwnedWordArt, "\n") {
		hkPut(g, 36, 16+i, line)
	}
	if p%10 < 2 {
		hkGlitch(g, int64(p)+13, 4)
	}
	return hkGreen(hkJoin(g))
}

const hkAccessArt = ` ███   ███   ███  █████  ████  ████
█   █ █   █ █   █ █     █     █
█   █ █     █     █     █     █
█████ █     █     ████   ███   ███
█   █ █     █     █         █     █
█   █ █   █ █   █ █         █     █
█   █  ███   ███  █████ ████  ████`

const hkGrantedArt = ` ███  ████   ███  █   █ █████ █████ ████
█   █ █   █ █   █ ██  █   █   █     █   █
█     █   █ █   █ █ █ █   █   █     █   █
█ ███ ████  █████ █  ██   █   ████  █   █
█   █ █ █   █   █ █   █   █   █     █   █
█   █ █  █  █   █ █   █   █   █     █   █
 ████ █   █ █   █ █   █   █   █████ ████`

const hkSystemArt = ` ████ █   █  ████ █████ █████ █   █
█     █   █ █       █   █     ██ ██
█      █ █  █       █   █     █ █ █
 ███    █    ███    █   ████  █ █ █
    █   █       █   █   █     █   █
    █   █       █   █   █     █   █
████    █   ████    █   █████ █   █`

const hkOwnedWordArt = ` ███  █   █ █   █ █████ ████
█   █ █   █ ██  █ █     █   █
█   █ █   █ █ █ █ █     █   █
█   █ █ █ █ █  ██ ████  █   █
█   █ █ █ █ █   █ █     █   █
█   █ ██ ██ █   █ █     █   █
 ███  █   █ █   █ █████ ████`

const hkSkullArt = `       ███████████
     ███████████████
   ███████████████████
   ███    █████    ███
  ████    █████    ████
   ███    █████    ███
   ████████   ████████
     █████     █████
      █           █
      █ ██ ██ ██ ██
       █         █
        █████████`
