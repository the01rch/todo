package main

import (
    . "fmt"
    . "os"
)

func main() {
    if !gestErr() {
        Exit(1)
    }
    //parseFile(getFile())
    //gestOpt()
    Printf("Best CLI todo app ever!\n")
}
