package main

import (
	. "encoding/json"
	. "fmt"
	"io/ioutil"
	. "os"
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
    if len(Args) == 1 {
        arr := json_to_array() 
        print_tasks(arr)
        return
    } else if !is_flag() {
        Println(help())
        Exit(1)
    }
    l := List{}
    arr := json_to_array()
    array2list(arr, l)
    tab := flag(arr, l)
    if tab == nil {
        Println("tab is nul")
    }
    test, err := Marshal(tab)
    if err != nil {
        Println("shit")
        return
    }
   err2 := ioutil.WriteFile("./data/list.json", []byte(test), 0777)
   if err2 != nil {
       Println("fdp")
   }
}
