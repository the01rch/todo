package main

import (
	. "os"
    . "strconv"
)

func flag(arr []data, l List) []data {
    switch Args[1] {
        case "-a":
            add_flag(l)
        case "-d":
            del_flag(l)
        default:
            arr = status_flag(arr)
    }
    return arr
}

func add_flag(l List) {
    tmp := l.head
    for i := 0; tmp.next != nil; tmp = tmp.next {
        str := Itoa(i)
        l.add(str, 0, Args[2])
        i++
    }  
}

func del_flag(l List) {
    for tmp := l.head; tmp.next != nil; tmp = tmp.next {
        if tmp.next.Id == Args[2] {
             tmp.next = tmp.next.next
        } 
    }
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

/*
func list2array(l List) string {
    tmp := l.head

    return ""
}
*/
