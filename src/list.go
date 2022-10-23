package main

import (
	. "encoding/json"
	. "fmt"
	. "os"
	"strconv"
)

type data struct {
    Id string
    Status int
    Name string
}

type Node struct {
    Id string `json:"Id"`
    Status int `json:"Status"`
    Name string `json:"Name"`
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
    check := 0
    for i := 0; i < len(arr); i++ {
        if arr[i].Status == 2 {
            check++
        }
    }
    Printf("\x1b[4m@Todo\033[0m ")
    Printf("[%d/%d]\n", check, len(arr))
    for i := 0; i < len(arr); i++ {
        Printf("  %s", arr[i].Id)
        if arr[i].Status == 0 {
            Printf(".   ")
        } else if arr[i].Status == 1 {
            Printf(".   ");
        } else if arr[i].Status == 2 {
            Printf(".   ");
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

func array2list(arr []data, l *List) {
    i := 0
    for ; i < len(arr); i++ {
        l.add(arr[i].Id, arr[i].Status, arr[i].Name)
    }
    l.len = i+1
}

func list2array(l *List) []data {
    tmp := l.head
    arr := make([]data, l.len)
    i := 0
    for ;tmp.next != nil; tmp = tmp.next {
        arr[i].Id = strconv.Itoa(i)
        arr[i].Name = tmp.Name
        arr[i].Status = tmp.Status
        i++
    }
    arr[i].Id = strconv.Itoa(i) 
    arr[i].Name = tmp.Name
    arr[i].Status = tmp.Status
    return arr
} 
