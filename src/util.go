package main

import (
    //"fmt"
    //"unsafe"
)

func is_alpha(c byte) bool {
    if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
        return true
    }
    return false
}

func strncmp(size int, str string, src string) bool {
    for i := 0; i < size; i++ {
        if src[i] != str[i] {
            return false
        }
    }
    return true
}

func strstr(str string, src string) bool {
    for i := 0; i < len(src); i++ {
        test := src[i:]
        if strncmp(len(str), str, test) == true {
            return true
        }
    }
    return false
}
