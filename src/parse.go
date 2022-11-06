package main

import (
    "fmt"
)

func count_rows(str []byte) int {
	rows := 0

	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			rows++
		}
	}
	return rows
}

func count_cols(str []byte, rows int) []int {
    cols := make([]int, rows)
    x := 0

    for y := 0; y < rows; y++ {
        for ; str[x] != '\n'; x++ {
            cols[y]++;
        }
        x++
    }
    return cols
}

func str_to_arr(str []byte) [][]byte {
    rows := count_rows(str)
    cols := count_cols(str, rows)
    arr := make([][]byte, rows)
    x := 0

    for y := 0; y < rows; y++ {
        arr[y] = make([]byte, cols[y])
        for w := 0; str[x] != '\n'; x++ {
            arr[y][w] = str[x]
            w++
        }
        x++
    }
    return arr
}

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

func is_alpha(c byte) bool {
    if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
        return true
    }
    return false
}

func gest_proj(str []byte) {
    arr := str_to_arr(str)

    fmt.Printf("\033[4mMy Boards :\033[0m\n")
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
