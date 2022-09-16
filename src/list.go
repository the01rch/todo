package main

import (
	. "os"
    . "fmt"
    . "encoding/json"
)

type data struct {
    Id string
    Status int
    Name string
}

type Node struct {
    Id string
    Status int
    Name string
    next *Node
} 

type List struct {
    head *Node
    len int
}

func json_to_array() []data {
    var arr []data

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

func print_tasks(arr []data) {
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

func (l *List) add(id string, status int, name string) {
    node := new(Node)
    node.Id = id
    node.Status = status
    node.Name = name
    if l.head == nil {
        l.head = node
    } else {
        tmp := l.head
        for ;tmp.next != nil; tmp = tmp.next {}
        tmp.next = node
    }
}

func (l List) print_list() {
    tmp := l.head
    for i := 0; tmp.next != nil; tmp = tmp.next {
        Printf("node = %d\n", i)
        Printf("id: %s, name : %s, status : %d\n", tmp.Id, tmp.Name, tmp.Status)
        i++
    }
}

func array2list(arr []data, l List) {
    for i := 0; i < len(arr); i++ {
        l.add(arr[i].Id, arr[i].Status, arr[i].Name)
    }
    l.print_list()
}

