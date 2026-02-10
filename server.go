package main

import (
    "bufio"
    "fmt"
    "net"
    "sync"
)

var (
    clients = make(map[net.Conn]bool)
    clientsMu sync.Mutex
)

func broadcast(message string, sender net.Conn) {
    clientsMu.Lock()
    defer clientsMu.Unlock()
    for client := range clients {
        if client != sender {
            fmt.Fprintln(client, message)
        }
    }
}

func handleConnection(conn net.Conn) {
    defer func() {
        clientsMu.Lock()
        delete(clients, conn)
        clientsMu.Unlock()
        conn.Close()
        broadcast("A client has disconnected.", conn)
        fmt.Println("Client disconnected:", conn.RemoteAddr())
    }()

    clientsMu.Lock()
    clients[conn] = true
    clientsMu.Unlock()
    broadcast("A new client has connected.", conn)
    fmt.Println("Client connected:", conn.RemoteAddr())

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        msg := scanner.Text()
        fmt.Printf("Received from %s: %s\n", conn.RemoteAddr(), msg)
        broadcast(fmt.Sprintf("%s says: %s", conn.RemoteAddr(), msg), conn)
    }
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server listening on port 8080...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}