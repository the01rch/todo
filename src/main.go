package main

import (
    . "os"
    . "fmt"
)

func help() string {
    return `
        $> ToDo [Flags] [Args]

        Flags:
            [-a] add a task ( a string argument is required ) 
            [-b] begin a task ( a task ID is required )
            [-c] check a task ( same as -b flag )
            [-d] delete a task ( same as -b flag )
            [-e] edit a task ( same as -b flag + another string argument )
        
        ToDo without Flags and Args display the list see the exemple bellow.

        Exemple:
            $> ToDo
                0 -   read the documentation
                1 -   write the README
                2 -   write the Makefile 
    `
}

func is_flag() bool {
    switch Args[1] {
        case "-a":
            return true
        case "-b":
            return true
        case "-c":
            return true
        case "-d":
            return true
        case "-e":
            return true
    }
    return false
}

func main() {
    if len(Args) > 1 && !is_flag() {
        Println(help())
        Exit(1)
    }
    Println(get_file())
}
