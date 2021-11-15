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

func main() {
    if !gestErr() {
        Exit(1)
    }
    Printf("Best CLI todo app ever!\n")
}
