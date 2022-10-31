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
    
    fmt.Println(rows)
    x := 0
    for y := 0; y < rows; y++ {
        for ; str[x] != '\n'; x++ {
            cols[y]++;
        }
        //fmt.Printf("cols[%d] : %d\n", y, cols[y])
        x++
    }
    return cols
}

func str_to_arr(str []byte) {
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
    for i := 0; i < rows; i++ { 
        fmt.Printf("arr[%d] : %s\n", i, string(arr[i]))
    }
}

/*
func get_pos(str []byte) []int {
    pos := make([]int, 2)
    check := 0

    for i := 0; i < len(str); i++ {
        if str[i] == '"' && check == 0 {
            check = 1
            pos[0] = i
        } else if str[i] == ']' && check == 1 {
            check = 2
            pos[1] = i
            break
        }
    } 
    return pos
}

func get_proj(str []byte) {
    rows := count_rows(str)
    //cols := count_cols(str)
    arr := make([][]byte, rows)

    x := 0
    for y := 0 ; y < rows; y++ {
        for ; str[x] != '\n'; x++ {
            arr[y][x] = str[x]
            fmt.Printf("arr[%d][%d] = %c\n", y, x, arr[y][x])
        }
        x++
    }
    
}

func get_size(str []byte) int {
    size := 0

    for i := 0; i < len(str); i++ {
        if str[i] == '"'
    }
    return size
}

func red(str []byte) []byte {
    i := 0
    arr := make([]byte, i)


    return arr 
}
*/
