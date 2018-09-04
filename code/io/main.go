package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("Missing file name argument")
        return
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    data := make([]byte, 100)
    spaces := 0
    for {
        data = data[:cap(data)]
        n, err := f.Read(data)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println(err)
            return
        }
        data = data[:n]
        for _, b := range data {
            if b == ' ' {
                spaces++
            }
        }
    }
    fmt.Println(spaces)
}