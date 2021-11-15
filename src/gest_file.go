package main

import (
	. "fmt"
	. "os"
	. "strconv"
)

func readFile() {
    content, err := ReadFile("./data/data_3.txt")
    isFile(err)
    parseFile(string(content))
    //Println(sliceArgs())
}

func isNum(c byte)bool {
    if c >= '0' && c <= '9' {
        return true
    }
    return false
}

func printTask(str []byte) {
    task := string(str)
    id := 0
    check := 0
    c := 0

    //Printf("\nThis is str[0] = %c\nThis is task = %s", str[0], task)
    if str[1] == '-' && isNum(str[2]) {
        id, _ = Atoi(string(task[2]))
        Printf("%d - ", id)
    }
    if isNum(str[0]) {
        check, _ = Atoi(string(task[0]))
        if check == 0 {
            Printf(" ")
        } else if check == 1 {
            Printf(" ")
        }
    }
    Printf("  \"")
    for i := 0; i < len(task); i++ {
        if str[i] == '\n' {
            c = 1
            i++
        }
        if c == 1 {
            Printf("%c", str[i])
        }
    }
    Printf("\"\n")
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
            //task = task[:0]
            //task = make([]byte, 100)
            x = 0
            if i+1 < len(str) {
                i++
            }
        }
        task[x] = str[i]
        x++
    }
}
