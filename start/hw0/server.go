package main

import (
    "fmt"
    "net"
    "os"
)

const (
    networkType    = "tcp"
    host           = "localhost"
    port           = "8080"
    serverAddress  = host + ":" + port
    successMessage = "OK\n" 
)

func main() {
    ln, err := net.Listen(networkType, serverAddress)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Ошибка запуска сервера: %v\n", err)
        os.Exit(1)
    }
    defer ln.Close() 

    fmt.Println("Сервер слушает на", serverAddress)

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Ошибка подключения: %v\n", err)
            continue 
        }

        go func(c net.Conn) {
            defer c.Close()
            _, err := c.Write([]byte(successMessage))
			if err != nil {
                fmt.Fprintf(os.Stderr, "Ошибка отправки ответа клиенту %s: %v\n", c.RemoteAddr(), err)
                return
            }
        }(conn)
    }
}