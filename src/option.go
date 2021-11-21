package main

import (
    . "fmt"
    . "os"
)

func helpOpt() {
    if Args[1] == "--help" {
        Printf("USAGE\n\ttodo [OPTIONS] task|ID\n\n")
        Printf("EXAMPLE\n\ttodo -a \"read a book\"\n")
        Printf("\ttodo -b [ID]\n\ttodo -c [ID]\n")
        Printf("\ttodo -d [ID]\n\ttodo -e [ID] \"read 2 books\"")
        Printf("\n\nOPTIONS\n\t-a\t\tadd a new task\n\t")
        Printf("-b\t\tbegin a task\n\t-c\t")
        Printf("\tcheck a task\n\t-d\t\tdelete a task\n\t-e\t")
        Printf("\tedit a task\n\t--help\tprint this help\n\n")
        Exit(0)
    }
}

func addOpt() {
    if Args[1] == "-a" {
        f, err := OpenFile("./data/data_3.txt", O_RDWR, 0644)
        isFile(err)
        n, err := f.Seek(0, SEEK_END)
        _, err = f.WriteAt([]byte("0-0\n"), n)
        n1, err := f.Seek(0, SEEK_END)
        _, err = f.WriteAt([]byte(Args[2]), n1)
        defer f.Close()
        Println("New task added !")
        Exit(0)
    }
}

func gestOpt() bool {
    if len(Args) == 1 {
        parseFile(getFile())
        Exit(0)
    }
    helpOpt()
    addOpt()
    return false
}
