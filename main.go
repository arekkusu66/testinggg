package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "time"
)


type Resp struct {
    Text string
    Ok   bool
}


type Client struct{
    Name string
    Resp chan Resp
}


func main() {
    var queuedUsers = make(chan Client, 1)

    go func() {
        for u := range queuedUsers {
            time.Sleep(time.Second * 6)
            u.Resp <- Resp{
                Text: strings.ToUpper(u.Name),
                Ok: true,
            }
        }
    }()

    var scanner = bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        var client = Client{
            Name: scanner.Text(),
            Resp: make(chan Resp, 1),
        }

        queuedUsers <- client

        go func(c Client) {
            var resp = <-c.Resp
            fmt.Println(resp.Ok, resp.Text)
        }(client)
    }
}
