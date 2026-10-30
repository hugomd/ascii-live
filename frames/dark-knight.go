package frame

import (
"fmt"
"time"
)

var dark-knight = FrameType{GetFrame: DefaultGetFrame(darkknightPR135Frames), GetLength: DefaultGetLength(darkknightPR135Frames), GetSleep: func() time.Duration { return 100 * time.Millisecond }}

var darkknightPR135Frames = []string{
"                 /\                 ",
"                /  \                ",
"       ********/    \********       ",
"      /                        \      ",
"     /   /\              /\     \     ",
"    /   /  \_______*****/  \     \    ",
"   /                            \   ",
"  /   /\                  /\     \  ",
" /***/  \**_______*******/  \*****\ ",
"       \                  /          ",
"        \**______________/           ",
}
