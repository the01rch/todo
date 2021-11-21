package main

import (
    . "fmt"
    . "os"
)

func isFile(e error) {
    if e != nil {
        panic(e)
        //Fprintf(Stderr, "File doesn't exist !\n")
        //Exit(1)
    }
}

func isOpt(str string) bool {
    switch {
        case str == "-a":
            return true
        case str == "-b":
            return true
        case str == "-c":
            return true
        case str == "-d":
            return true
        case str == "-e":
            return true
        case str == "--help":
            return true
    }
    Fprintf(Stderr, "Wrong option format, try --help for help.\n")
    return false
}

func gestErr() bool {
    if !gestOpt() && !isOpt(Args[1]) {
        return false
    }
    return true
}
