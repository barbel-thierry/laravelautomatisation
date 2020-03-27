package main

import (
    "fmt"
    "os"
    "laravelnew"
    "laravelreset"
    "laravelopen"
)

func main() {
    cmd := os.Args[1]
    name := ""
    if len(os.Args) > 2 {
        name = os.Args[2]
    }

    switch cmd {
    case "new":
        laravelnew.Project(name)
    case "reset":
        laravelreset.CreateOrReset()
    case "open":
        laravelopen.Run()
    case "setup":
    default:
        fmt.Println("Sorry, but this command doesn't exist.")
    }
}