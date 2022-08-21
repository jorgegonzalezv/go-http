package main 

import (
    "fmt"
    "net"
    "os"
)

// server configuration 
const (
    connHost = "0.0.0.0"
    connPort = "8080"
    connType = "tcp"
)

func main() {
    fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)

    // listen to tcp socket at given port number
    listener, err := net.Listen(connType, connHost+":"+connPort)

    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer listener.Close()

    // client accept request loop
    for {
        c, err := listener.Accept() // blocking call
        if err != nil {
            fmt.Println("Error connecting:", err.Error())
            return
        }
        fmt.Println("Client connected.")

        fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
    }   
}
