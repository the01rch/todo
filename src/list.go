package main

import (
	. "os"
    . "fmt"
    . "encoding/json"
)

type Node struct {
    Id string
    Status int
    Name string
}

func array_to_json(arr []Node) {
    json, _ := Marshal(arr)
    Println(json)
}

func json_to_array() []Node {
    var arr []Node

    data, err := ReadFile("./data/list.json")
    if err != nil {
        Println(err)
    }
    err2 := Unmarshal(data, &arr)
    if err2 != nil {
        Println(err2)
    } 
    return arr
}

func print_tasks(arr []Node) {
    for i := 0; i < len(arr); i++ {
        Printf("%s", arr[i].Id)
        if arr[i].Status == 0 {
            Printf(" -   ")
        } else if arr[i].Status == 1 {
            Printf(" -   ");
        } else if arr[i].Status == 2 {
            Printf(" -   ");
        }
        Printf("%s\n", arr[i].Name)
    }
}

func array_to_list(arr []Node) {
    
}
