package main

import (
    . "fmt"
)

func update_status(arr []Node, id string, status int) {
    for i := 0; i < len(arr); i++ {
        if status == BEGIN && arr[i].Id == id {
            arr[i].Status = 1
            break
        }
    }
    Println(arr)
}
