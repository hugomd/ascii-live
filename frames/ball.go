package main

import (
    "fmt"
    "time"
)

var frames = []string{
    `
o
    `,
    `
 o
    `,
    `
  o
    `,
    `
   o
    `,
    `
  o
    `,
    `
 o
    `,
    `
o
    `,
}


func main() {
    for {
        for _, f := range frames {
            fmt.Print("\033[H\033[2J") // clear screen
            fmt.Println(f)
            time.Sleep(100 * time.Millisecond)
        }
    }
}
