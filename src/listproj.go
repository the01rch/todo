package main

import (
    "fmt"
)

func infos_proj(arr [][]byte) []int {
    nb := 0
    for y := 0; y < len(arr); y++ {
        for x := 0; x < len(arr[y]); x++ {
            if arr[y][x] == '[' {
                nb++
            }
        }
    }
    infos := make([]int, nb+1)
    infos[0] = nb
    i := 1
    for y := 0; y < len(arr); y++ {
        for x := 0; x < len(arr[y]); x++ {
            if arr[y][x] == '[' {
                infos[i] = y
                i++
            }
        }
    }
    return infos
}

func proj_names(arr [][]byte) [][]byte {
    infos := infos_proj(arr)
    proj_names := make([][]byte, infos[0])
    index := 0

    for i := 1; i < infos[0]+1; i++ {
        for y := 0; y < len(arr); y++ {
            if infos[i] == y {
                proj_names[index] = make([]byte, len(arr[y])) 
                proj_names[index] = arr[y]
                index++
            }
        }
    }
    return proj_names
}

func gest_proj(str []byte) {
    arr := str_to_arr(str)

    fmt.Printf("\033[4mBoards :\033[0m\n")
    projn := proj_names(arr)
    for y := 0; y < len(projn); y++ {
        fmt.Printf(" @")
        for x := 0; x < len(projn[y]); x++ {
            if is_alpha(projn[y][x]) {
                fmt.Printf("%c", projn[y][x])
            }
        }
        fmt.Printf("\n");
    }
}
