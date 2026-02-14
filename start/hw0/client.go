package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const (
    networkType   = "tcp"
    host          = "localhost"
    port          = "8080" 
    serverAddress = host + ":" + port
    readTimeout   = 5 * time.Second
    successMessage = "OK\n"
)

func main() {
    conn, err := net.Dial(networkType, serverAddress)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Ошибка подключения к %s: %v\n", serverAddress, err)
        os.Exit(1)
    }
    defer conn.Close()

    conn.SetReadDeadline(time.Now().Add(readTimeout))

    _, err = bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Fprintf(os.Stderr, "Ошибка чтения данных: %v\n", err)
        return
    }

    fmt.Print(successMessage)
}