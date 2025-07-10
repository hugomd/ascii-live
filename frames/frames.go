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
	"minecraft":       Minecraft,
	"earth":           Earth,
	"kitty":           Kitty,
	"india":           India,
	"brittany":        Brittany,
}
