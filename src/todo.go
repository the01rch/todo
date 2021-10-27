package main

import (
    . "fmt"
    . "os"
)

func sliceArgs()[][]string {
    size := len(Args)
    arr := make([][]string, size)

    for i := 0; i < size; i++ {
        arr[i] = []string{Args[i]}
    }
    return arr
}

func readFile() {
    _, err := ReadFile("./data/data_3.txt")
    isFile(err)
    Println(sliceArgs())
}

func helpOption() {
    Printf("USAGE: todo [OPTIONS] task|ID\n\t\tExample: todo -a 'read a book'")
    Printf("\n\nOPTIONS\n\t-b, --begin\t\t\tbegin a task\n\t-c, --check\t\t")
    Printf("\tcheck a task\n\t-d, --delete\t\t\tdelete a task\n\t-e, --edit\t")
    Printf("\t\tedit a task\n\t -h, --help\t\t\tprint this help\n")
}

func isOption()bool {
    if len(Args) == 1 {
        readFile()
        Exit(0)
    }
    if Args[1] == "--h" || Args[1] == "--help" {
        helpOption()
        Exit(0)
    }
    return false
}

func main() {
    if !gestErr() {
        Exit(1)
    }
    readFile()
}
