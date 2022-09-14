package main

import (
    . "os"
    . "fmt"
)

const (
    BEGIN = 1
    CHECK = 2
    EDIT = 3
)

func help() string {
    return `
        $> ToDo [Flags] [Args]

        Flags:
            [-a] add a task ( a string argument is required ) 
            [-b] begin a task ( a task ID is required )
            [-c] check a task ( same as -b flag )
            [-d] delete a task ( same as -b flag )
            [-e] edit a task ( same as -b flag + string argument )
        
        ToDo without Flags and Args display the list see the exemple bellow.

        Exemple:
            $> ToDo
                0 -   read the documentation
                1 -   write the README
                2 -   write the Makefile 
    `
}

func is_flag() bool {
    if len(Args) < 3 {
        return false
    }
    switch {
        case Args[1] == "-a" && len(Args) <= 4:
            return true
        case Args[1] == "-b":
            return true
        case Args[1] == "-c":
            return true
        case Args[1] == "-d":
            return true
        case Args[1] == "-e" && len(Args) == 4:
            return true
    }
    return false
}

func main() {
    if len(Args) > 1 && !is_flag() {
        Println(help())
        Exit(1)
    }
    arr := json_to_array() 
    if len(Args) == 1 {
        print_tasks(arr)
        return
    }
    if len(Args) > 2 && Args[1] == "-b" {
        update_status(arr, Args[2], BEGIN)
    }
    if len(Args) > 2 && Args[1] == "-c" {
        update_status(arr, Args[2], CHECK)
    }
    if len(Args) > 2 && Args[1] == "-d" {
        update_status(arr, Args[2], BEGIN)
    }
    if len(Args) > 2 && Args[1] == "-e" {
        update_status(arr, Args[2], BEGIN)
    }
}
