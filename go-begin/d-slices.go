package main

import (
    "fmt"
)

func slice() {
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
    fmt.Println(slice)

    // New slice out of old slice
    slice2 := slice[1:]
    fmt.Println(len(slice2))
    fmt.Println(cap(slice2))
    fmt.Println(slice2)

    // Modifying new slice will modify old slice also
    slice2[1]=22
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
    fmt.Println(slice)
    fmt.Println(len(slice2))
    fmt.Println(cap(slice2))
    fmt.Println(slice2)
}
