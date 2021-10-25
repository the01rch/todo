package main

import (
    "fmt"
    "os"
    //"bufio"
)

func isArgs() bool {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Too few arguments !\n")
        return false
    }
    return true
}

func isFile(e error) {
    if e != nil {
        fmt.Fprintf(os.Stderr, "File doesn't exist !\n")
        os.Exit(1)
    }
}

func sliceArgs()[][]string {
    size := len(os.Args)
    arr := make([][]string, size)

    for i := 0; i < size; i++ {
        arr[i] = []string{os.Args[i]}
    }
    return arr
}

func isOption(str string) bool {
    switch {
        case str == "add":
            return true
        case str == "rem":
            return true
        case str == "list":
            return true
        case str == "check":
            return true
        case str == "help":
            return true
    }
    return false
}

func main() {
    if !isArgs() {
        os.Exit(1)
    }
    _, err := os.ReadFile("./data/data_3.txt")
    isFile(err)
    fmt.Println(sliceArgs())
}
