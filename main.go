package main

import (
    "fmt"
    "encoding/hex"
)

func main() {
    fmt.Println(hex.EncodeToString([]byte("hello!")))
}
