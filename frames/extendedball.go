package frames

import "time"

var Ball = FrameType{GetFrame: DefaultGetFrame(ballPR127Frames), GetLength: DefaultGetLength(ballPR127Frames), GetSleep: func() time.Duration { return 100 * time.Millisecond }}

var ballPR127Frames = []string{
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
