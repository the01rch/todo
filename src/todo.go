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
    content, err := ReadFile("./data/data_3.txt")
    isFile(err)
    parseFile(string(content))
    //Println(sliceArgs())
}

func printTask(str []byte) {
    task := string(str)
    Printf("%s", task)
    Println("\nTHAT A TASSK !")
}

func parseFile(str string) {
    task := make([]byte, 100)
    y := 0
    x := 0

    for i := 0; i < len(str); i++ {
        if str[i] == '\n' {
            y++
        }
        if y == 2 {
            y = 0
            printTask(task)
            x = 0
        }
        task[x] = str[i]
        x++
    }
    return
}

func main() {
    if !gestErr() {
        Exit(1)
    }
    Printf("Best CLI todo app ever!\n")
}
