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
    next *Node
}

type List struct {
    head *Node
    len int
}

/*
func (l *List) insert(name string, status int, id string) {
    n := Node{}
    n.Name = name
    n.Status = status
    n.Id = id
    if l.len == 0 {
        l.head = &n
        l.len++
        return
    }
    ptr := l.head
    for i := 0; i < l.len; i++ {
        if ptr.next == nil {
            ptr.next = &n
            l.len++
            return
        }
        ptr = ptr.next
    }
}

func (l *List) Print() {
    if l.len == 0 {
        Println("No nodes in the list !")
        return
    }
    ptr := l.head
    for i := 0; i < l.len; i++ {
        Printf("%s - ", ptr.Id)
        switch ptr.Status {
        case 0:
            Printf("  ")
        case 1:
            Printf("  ")
        case 2:
            Printf("  ")
        }
        Println(ptr.Name)
        ptr = ptr.next
    }
}

func print_array(arr []Node) {
    for i := 0; i < len(arr); i++ {
        s := arr[i]
        Printf("%s - ", s.Id)
        switch s.Status {
        case 0:
            Printf("  ")
        case 1:
            Printf("  ")
        case 2:
            Printf("  ")
        }
        Printf("%s\n", s.Name)
    }
}

func array_to_list(arr []Node) {
    list := List{}
    for i := 0; i < len(arr); i++ {
        s := arr[i]
        list.insert(s.Name, s.Status, s.Id)
    }
}

*/

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
