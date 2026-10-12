package frames

import "time"

var ExtendedBall = FrameType{GetFrame: DefaultGetFrame(extendedBallFrames), GetLength: DefaultGetLength(extendedBallFrames), GetSleep: func() time.Duration { return 100 * time.Millisecond }}

var extendedBallFrames = []string{
	"\no\n    ",
	"\n o\n    ",
	"\n  o\n    ",
	"\n   o\n    ",
	"\n    o\n    ",
	"\n     o\n    ",
	"\n      o\n    ",
	"\n       o\n    ",
	"\n        o\n    ",
	"\n         o\n    ",
	"\n          o\n    ",
	"\n           o\n    ",
	"\n            o\n    ",
	"\n             o\n    ",
	"\n              o\n    ",
	"\n             o\n    ",
	"\n            o\n    ",
	"\n           o\n    ",
	"\n          o\n    ",
	"\n         o\n    ",
	"\n        o\n    ",
	"\n       o\n    ",
	"\n      o\n    ",
	"\n     o\n    ",
	"\n    o\n    ",
	"\n   o\n    ",
	"\n  o\n    ",
	"\n o\n    ",
}
