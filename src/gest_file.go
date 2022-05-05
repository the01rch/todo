package main

import (
	. "os"
    . "fmt"
)

func is_file_error(e error) {
    if e != nil {
        Println("The list.json file is not readable !")
        Exit(1)
    }
}

func get_file() string {
    content, err := ReadFile("./data/list.json")
    is_file_error(err)
    str := string(content)
    return str
}

func is_num(c byte)bool {
    if c >= '0' && c <= '9' {
        return true
    }
    return false
}
