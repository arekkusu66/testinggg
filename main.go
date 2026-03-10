package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string, 1)

    go func() {
        time.Sleep(time.Second * 6)
        ch <- "jjj"
    }()

    fmt.Println(<-ch)
}
