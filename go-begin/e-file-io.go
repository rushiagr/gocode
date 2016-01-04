package main

import (
    "fmt"
    "io/ioutil"
)

func fileio() {
    data := []byte("some temp data")
    fmt.Println(data)
    err := ioutil.WriteFile("/tmp/gowrite", data, 0644)
    //err := ioutil.WriteFile("/root/gowrite", data, 0644)
    if err != nil {
        panic(err)
    }

    data, err = ioutil.ReadFile("/tmp/gowrite")
    fmt.Println(data)
    fmt.Println(string(data))
}
