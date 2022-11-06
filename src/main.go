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

func is_proj_syntax(proj []byte) bool {
    if proj[0] != '@' {
        fmt.Printf("Wrong syntax ! try --help\n") 
        return false
    }
    return true
}

func strncmp(index int, str1 []byte, str2 []byte) bool {
    y := 0

    for i := index; i < len(str1); i++ {
        if str1[i] != str2[y] {
            return false
        }
        y++
    }
    return true
}


func main() {
	if len(os.Args) == 1 {
		json_file := get_file()
        gest_proj(json_file)
		return
	} else if len(os.Args) == 2 {
        proj := os.Args[1]

        if !is_proj_syntax(proj) {
            os.Exit(1)
        } 
         
        //arr := json_to_array()
		//fmt.Println(arr)
		//print_tasks(arr)
        return
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
