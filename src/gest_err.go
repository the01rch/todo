package main

import (
    . "fmt"
    . "os"
)

func isFile(e error) {
    if e != nil {
        Fprintf(Stderr, "File doesn't exist !\n")
        Exit(1)
    }
}

func isArgs() bool {
    if len(Args) < 1 {
        Fprintf(Stderr, "Too few arguments !\n")
        return false
    }
    return true
}

func gestOption(str string) bool {
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
    }
    Fprintf(Stderr, "Wrong option, try -h for help.\n")
    return false
}

func gestErr() bool {
    if !isArgs() {
        return false
    }
    if !isOption() && !gestOption(Args[1]) {
        return false
    }
    return true
}
