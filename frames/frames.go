package frames

import "time"

type FrameType struct {
	GetFrame  func(int) string
	GetLength func() int
	GetSleep  func() time.Duration
}

// Create a function that returns the next frame, based on length
func DefaultGetFrame(frames []string) func(int) string {
	return func(i int) string {
		return frames[i%(len(frames))]
	}
}

// Create a function that returns frame length
func DefaultGetLength(frames []string) func() int {
	return func() int {
		return len(frames)
	}
}

// Sleep time between frames
func DefaultGetSleep() func() time.Duration {
	return func() time.Duration {
		return time.Millisecond * 70
	}
}

// Given frames, create a FrameType with those frames
func DefaultFrameType(frames []string) FrameType {
	return FrameType{
		GetFrame:  DefaultGetFrame(frames),
		GetLength: DefaultGetLength(frames),
		GetSleep:  DefaultGetSleep(),
	}
}

var FrameMap = map[string]FrameType{
	"airplane":        Airplane,
	"globe":           Globe,
	"batman":          Batman,
	"batman-running":  BNR,
	"bnr":             BNR,
	"can-you-hear-me": Rick,
	"clock":           Clock,
	"coin":            Coin,
	"donut":           Donut,
	"dvd":             Dvd,
	"forrest":         Forrest,
	"hes":             HES,
	"knot":            TorusKnot,
	"nyan":            Nyan,
	"parrot":          Parrot,
	"rick":            Rick,
	"spidyswing":      Spidy,
	"torus-knot":      TorusKnot,
	"purdue":          Purdue,
	"as":              AStrend,
	"bomb":            Bomb,
	"maxwell":         Maxwell,
	"earth":           Earth,
	"kitty":           Kitty,
	"india":           India,
	"brittany":        Brittany,
	"yasir":           Yasir,
	"pig":             Pig,
	"bodypack":        Bodypack,
	"floss":           Floss,
	"ring":            Ring,
	"fourtwo":         FourTwo,
	"Furry":           Furry,
	"simpson-bush":    Simpson,
	"skull":           Skull,
	"oiia-cat":        Oiia,
	"forty-two":       FortyTwo,
	"duck":            Duck,
	"ethel":           Ethel,
	"gagner":          Hello,
	"perdu":           Hello,
	"hello":           Hello,
	"dino":            Dino,
	"tesseract":       Tesseract,
	"happybirthday":   HappyBirthday,
	"blahaj":          Blahaj,
	"gopher":          Gopher,
	"teddy":           Teddy,
	"fish":            Fish,
	"hack":            Hack,
	"rocket":          Rocket,
	"mewrizz":         Mewrizz,
	"test":            Test,
	"chomik":          Chomik,
	"jd":              JD,
	"macarenavisbal":  Macarena,
	"badapple":        BadApple,
	"minecraft":       Minecraft,
	"sweaty-gamer":    SweatyGamer,
	"gandalf":         Gandalf,
	"eclipseos":       EclipseOS,
	"dozer":           Dozer,
	"pikachu":         Pikachu,
	"hatsunemiku":     Miku,
	"gon":             Gon,
	"cjhhaonnonuyr":   Cjhhaonnonuyr,
	"pennywise":       Pennywise,
	"maya":            AbejaMaya,
	"SodaPop":         SodaPop,
	"spaceship":       Spaceship,
	"shrek":           Shrek,
	"matrix":          Matrix,
	"ball":            Ball,
	"extendedball":    ExtendedBall,
	"kke":             Kke,
	"dancingcoin":     DancingCoin,
	"anonymous":       Anonymous,
	"analogclock":     AnalogClock,
	"drink":           Drink,
	"blueshirt":       Blueshirt,
	"heart":           Heart,
	"frisk":           FriskDance,
	"sans":            Sans10thanniversary,
	"sixseven":        SixSeven,
	"pedro":           Pedro,
	"bangladesh":      Bangladesh,
	"hell":            Hell,
	"gandalf-sax":     GandalfSax,
	"egg":             Egg,
	"moonwalk":        Moonwalk,
	"rainbow-tunnel":  RainbowTunnel,
	"aurora":          Aurora,
	"plasma":          Plasma,
	"helix":           Helix,
	"cube-lite":       CubeLite,
	"luv":             Luv,
}
