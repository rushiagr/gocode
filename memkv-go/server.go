package main

import (
    "fmt"
    "bufio"
    "net"
    "strings"
)

// Copied from https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

func server() {
    fmt.Println("Starting server...")

    ln, _ := net.Listen("tcp", ":8080")
    conn, _ := ln.Accept()

    for {
        msg, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println("Message received:", string(msg))
        newmsg := strings.ToUpper(msg)
        conn.Write([]byte(newmsg + "\n"))
    }
}
