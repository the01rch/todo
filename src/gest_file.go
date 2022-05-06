package main

import (
	. "os"
    . "fmt"
    . "encoding/json"
)

type List struct {
    Name string
    Id int
    Status int
}

/*
type List struct {
    head *Node
    len int
}

func (l *List) Insert(name string, status int, id int) {
    n := Node{}
    n.name = name
    n.status = status
    n.id = id
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
        Printf("%d - ", ptr.id)
        switch ptr.status {
        case 0:
            Printf("  ")
        case 1:
            Printf("  ")
        case 2:
            Printf("  ")
        }
        Println(ptr.name)
        ptr = ptr.next
    }
}
*/

func print_list(arr []List) {
    s := arr[0]
    for i := 0; i < len(arr); i++ {
        s = arr[i]
        Printf("%d - ", s.Id)
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

func json_to_struct() {
    var arr []List

    data, err := ReadFile("./data/list.json")
    if err != nil {
        Println(err)
    }
    err2 := Unmarshal(data, &arr)
    if err2 != nil {
        Println(err2)
    } 
    print_list(arr)
}
