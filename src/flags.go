package main

import (
    "os"
    "strconv"
)

func flag(arr []data, l *List) []data {
    switch os.Args[1] {
        case "-a":
            add_flag(l)
            arr = list2array(l)
        case "-d":
            del_flag(l)
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
    node.Name = strconv.Itoa(i)
    node.Status = 0
    node.Name = os.Args[2]
    tmp.next = node
}

func del_flag(l *List) {
    i := 0
    taille := len(os.Args)
    check := 0
    for y := 2; y < taille; y++ { 
        for tmp := l.head; tmp.next != nil; tmp = tmp.next {
            if i == 0 && os.Args[y] == "0" {
                tmp = tmp.next
                l.head = tmp
                check++
                break
            } else if tmp.next.Id == os.Args[y] && tmp.next.next == nil {
                check++
                tmp.next = nil 
                break
            } else if tmp.next.Id == os.Args[y] {
                check++
                tmp.next = tmp.next.next
            } 
            i++
        }
    }
    l.len = l.len - (check + 1)
}

func edit_proj_name() {
    _, _, _, a := board_proj()
    projn := ""
    projn = os.Args[3]
    projn = projn[1:]
    s := strconc([]byte("\""), []byte(os.Args[2]))
    for y := 0; y < len(a); y++ {
        if strstr(projn, string(a[y])) {
            s = strconc(s, []byte("\""))
            s = strconc(s, []byte(": ["))
            a[y] = []byte(s)
            break
        }
    }
    home := os.Getenv("HOME")
    os.WriteFile((home + "/todo.json"), arr2str(a,0,len(a)), 0777)
}

func status_flag(arr []data) []data {
    taille := len(os.Args)
    _, err := strconv.Atoi(os.Args[2])
    if os.Args[1] == "-e" && err != nil {
            edit_proj_name()
            os.Exit(0)
    }
    for i := 0; i < len(arr); i++ {
        for y := 2; y < taille; y++ {
            if os.Args[y] == arr[i].Id {
                switch os.Args[1] {
                case "-b":
                    arr[i].Status = BEGIN
                case "-c":
                    arr[i].Status = CHECK
                case "-e":
                    arr[i].Name = os.Args[3]
                }
            }
        }
    }
    return arr
}
