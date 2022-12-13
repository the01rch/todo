package main

import (
	"fmt"
	"os"
	// "io/ioutil")
)

const (
	BEGIN = 1
	CHECK = 2
	EDIT  = 3
)

func help() string {
	return `
        todo [Flags] [Args]
    `
}

func is_flag() bool {
	if len(os.Args) < 3 {
		return false
	}
	switch {
	case os.Args[1] == "-a" && len(os.Args) <= 4:
		return true
	case os.Args[1] == "-b":
		return true
	case os.Args[1] == "-c":
		return true
	case os.Args[1] == "-d":
		return true
	case os.Args[1] == "-e" && len(os.Args) == 4:
		return true
	}
	return false
}

func is_proj_syntax(proj string) bool {
    if proj[0] != '@' {
        fmt.Printf("Wrong syntax ! try --help\n") 
        return false
    }
    return true
}

func is_proj_name(projn [][]byte, projA string) bool {
    exist := false 
    for y := 0; y < len(projn); y++ {
        if strstr(projA, string(projn[y])) == true {
            exist = true
            return true
        }
    }
    if !exist {
        fmt.Printf("comment sa mon boeuf ?\n")
    }
    return false
}

func index_end_proj(arr [][]byte, begin int) int {
    for y := begin; y < len(arr); y++ {
        for x := 0; x < len(arr[y]); x++ {
            if arr[y][x] == ']' {
                return y
            }
        }
    }
    return 0
}

func index_begin_proj(arr [][]byte, projA string) int {
    begin := 0
    infos := infos_proj(arr)
    for i := 1; i < len(infos); i++ {
        if strstr(projA, string(arr[infos[i]])) {
            begin = infos[i] + 1
        }
    }
    return begin
} 

func size_proj_select(begin int, end int, arr [][]byte) int {
    size := 0

    for y := begin; y < end; y++ {
        for x := 0; x < len(arr[y]); x++ {
            size++
        } 
    } 
    return size
}

func select_proj(size int, arr [][]byte, begin int, end int) []byte {
    diff := end - begin
    p := make([]byte, size + diff + 3)
    w := 2

    p[0] = '['
    p[1] = '\n'
    for y := begin; y < end; y++ {
        for x := 0; x < len(arr[y]); x++ {
            p[w] = arr[y][x]
            w++
        }
        p[w] = '\n'
        w++
    }
    p[w] = ']'
    return p
}

func board_proj() {
    json_file := get_file()
    arr := str_to_arr(json_file) 
    projn := proj_names(arr)
    projA := os.Args[1]
    if !is_proj_syntax(projA) {
        os.Exit(1)
    } 
    projA = projA[1:]
    if !is_proj_name(projn, projA) {
        os.Exit(1)
    }
    begin := index_begin_proj(arr, projA)
    end := index_end_proj(arr, begin)
    size := size_proj_select(begin, end, arr)
    p := select_proj(size, arr, begin, end)
    //fmt.Printf("%s", p)
    tab := json_to_array(p)
    print_tasks(tab)
}

func main() {
	if len(os.Args) == 1 {
		json_file := get_file()
        gest_proj(json_file)
		return
	} else if len(os.Args) == 2 {
        board_proj()
    } else if !is_flag() {
        fmt.Println(help())
        os.Exit(1)
    }
    /*
	   l := List{}
	   arr := json_to_array()
	   array2list(arr, &l)
	   tab := flag(arr, &l)
	   test, _ := MarshalIndent(tab, "", " ")
	   ioutil.WriteFile("./data/list.json", []byte(test), 0777)
	*/
}
