package main

import (
	. "os"
)

func getFile() string {
    content, err := ReadFile("./data/data_3.txt")
    is_file(err)
    str := string(content)
    return str
}

func isNum(c byte)bool {
    if c >= '0' && c <= '9' {
        return true
    }
    return false
}
