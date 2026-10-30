package main

import (
"fmt"
"time"
)

func main() {
batman := []string{
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

```
fmt.Print("\033[2J\033[H")

for _, line := range batman {
	fmt.Println(line)
	time.Sleep(150 * time.Millisecond)
}

fmt.Println("\n          THE DARK KNIGHT")
```

}
