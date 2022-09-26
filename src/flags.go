package main

import (
	. "os"
    . "strconv"
)

func flag(arr []data, l *List) []data {
    switch Args[1] {
        case "-a":
            add_flag(l)
            arr = list2array(l)
        case "-d":
            if len(Args) > 4 {
                for i := 2; i < len(Args); i++ {
                    del_flag(l, i)
                }
            }
            arr = list2array(l)
        default:
            arr = status_flag(arr)
    }
    return arr
}

func add_flag(l *List) {
    tmp := l.head
    l.head = tmp
    i := 0
    for ; tmp.next != nil; tmp = tmp.next {
        i++
    }  
    node := new(Node)
    node.Id = Itoa(i)
    node.Status = 0
    node.Name = Args[2]
    tmp.next = node
}

func del_flag(l *List, index int) {
    i := 0
    for tmp := l.head; tmp.next != nil; tmp = tmp.next {
        if i == 0 && Args[index] == "0" {
            tmp = tmp.next
            l.head = tmp
            break
        } else if tmp.next.Id == Args[index] && tmp.next.next == nil {
            tmp.next = nil 
            break
        } else if tmp.next.Id == Args[index] {
            tmp.next = tmp.next.next
        } 
        i++
    }
    l.len = l.len - 2
}

func status_flag(arr []data) []data {
    for i := 0; i < len(arr); i++ {
        if Args[2] == arr[i].Id {
            switch Args[1] {
                case "-b":
                    arr[i].Status = BEGIN
                case "-c":
                    arr[i].Status = CHECK
                case "-e":
                    arr[i].Name = Args[3]
            }
        }
    }
    return arr
}
