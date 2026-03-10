package main

import (
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

var q = `CREATE TABLE users (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE CHECK(username <> ''),
    join_date DATETIME NOT NULL DEFAUL DATETIME('now')
)`

func main() {
    db, err := sql.Open("sqlite3", "mydb.db")

    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    if _, err := db.Exec(q); err != nil {
        fmt.Println(err)
        return
    }
}
