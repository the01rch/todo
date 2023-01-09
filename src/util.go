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

func get_size(arr [][]byte, begin int, end int) int {
    taille := 0    
    for y := begin; y < end; y++ {
        for x := 0; x < len(arr[y]); x++ {
            taille++
        }
        taille++
    }
    return taille
}

func arr2str(arr [][]byte, begin int, end int) []byte {
    tab := make([]byte, get_size(arr, begin, end))

    for y, c := begin, 0; y < end; y++ {
        for x := 0; x < len(arr[y]); x++ {
            tab[c] = arr[y][x]
            c++
        }
        tab[c] = '\n'
        c++
    }
    return tab
}

func strconc(s1 []byte, s2 []byte) []byte {
    taille := len(s1) + len(s2)
    str := make([]byte, taille)

    for i, c := 0, 0; i < taille; i++ {
        if i < len(s1) {
            str[i] = s1[i]
        } else {
            str[i] = s2[c]
            c++
        }
    }
    return str
}
