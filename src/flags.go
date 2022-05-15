package main

import (
    . "fmt"
)

func begin_flag(arr []Node, id string) {
    for i:=0; i < len(arr); i++ {
        if arr[i].Id == id {
            arr[i].Status = 1
            break
        }
    }
    Println(arr)
}
