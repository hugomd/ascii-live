package frames

type FrameType struct {
	GetFrame  func(int) string
	GetLength func() int
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

// Given frames, create a FrameType with those frames
func DefaultFrameType(frames []string) FrameType {
	return FrameType{
		GetFrame:  DefaultGetFrame(frames),
		GetLength: DefaultGetLength(frames),
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
	"forrest":         Forrest,
	"hes":             HES,
	"knot":            TorusKnot,
	"nyan":            Nyan,
	"parrot":          Parrot,
	"playstation":     PlayStation,
	"rick":            Rick,
	"spidyswing":      Spidy,
	"torus-knot":      TorusKnot,
}
