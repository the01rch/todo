package main

import (
//	"fmt"
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

func main() {
	if len(os.Args) == 1 {
		arr := json_to_array()
        gest_proj(arr)
		//fmt.Println(arr)
		//print_tasks(arr)
		return
	}
	/*
	   else if !is_flag() {
	       Println(help())
	       Exit(1)
	   }
	   l := List{}
	   arr := json_to_array()
	   array2list(arr, &l)
	   tab := flag(arr, &l)
	   test, _ := MarshalIndent(tab, "", " ")
	   ioutil.WriteFile("./data/list.json", []byte(test), 0777)
	*/
}
