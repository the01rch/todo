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


func count_cols(str []byte) []int {
    rows := count_rows(str)
    cols := make([]int, rows)
    
    fmt.Println(rows)
    x := 0
    for y := 0; y < rows; y++ {
        for ; str[x] != '\n'; x++ {
            cols[y]++;
        }
        x++
    }
    return cols
}

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

func get_proj(str []byte)  {
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
