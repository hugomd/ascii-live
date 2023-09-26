package frames

type FrameType struct {
	GetFrame  func(int) string
	GetLength func() int
}

// Create a function that returns the next frame, based on length
func DefaultGetFrame(frames []string) func(int) string {
	return func(i int) string {
		return frames[i%(len(frames)-1)]
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
	"forrest":         Forrest,
	"parrot":          Parrot,
	"clock":           Clock,
	"nyan":            Nyan,
	"rick":            Rick,
	"can-you-hear-me": Rick,
	"donut":           Donut,
	"blink":	   Blink,
}
