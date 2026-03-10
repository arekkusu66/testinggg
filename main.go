package main

import (
    "fmt"
    "encoding/hex"
)

func main() {
    ch := makr(chan string, 1)

    go func() {
        time.Sleep(time.Second * 6)
        ch <- "jjj"
    }()

    fmt.Println(<-ch)
}
