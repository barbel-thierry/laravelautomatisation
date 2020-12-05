package main

import (
    "fmt"
    "os"
    "github.com/barbel-thierry/laravelautomatisation/laravelnew"
    "github.com/barbel-thierry/laravelautomatisation/laravelreset"
    "github.com/barbel-thierry/laravelautomatisation/laravelopen"
)

func main() {
    cmd := os.Args[1]
    name := ""
    version := ""

    if len(os.Args) > 2 {
        name = os.Args[2]
    }

    if len(os.Args) > 3 {
        version = os.Args[3]
    }

    switch cmd {
    case "new":
        laravelnew.Project(name, version)
    case "reset":
        laravelreset.CreateOrReset(name)
    case "open":
        laravelopen.Run()
    case "setup":
    default:
        fmt.Println("Sorry, but this command doesn't exist.")
    }
}
