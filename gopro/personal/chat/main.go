package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "sync"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

// Define a structure to hold connected clients
var clients = make(map[*websocket.Conn]bool)
var mutex = &sync.Mutex{}

// broadcast sends a message to all connected clients
func broadcast(message []byte) {
    mutex.Lock()
    defer mutex.Unlock()
    for client := range clients {
        err := client.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            log.Printf("write error: %v", err)
            client.Close()
            delete(clients, client)
        }
    }
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    // Register the new client
    mutex.Lock()
    clients[conn] = true
    mutex.Unlock()

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("read error:", err)
            mutex.Lock()
            delete(clients, conn)
            mutex.Unlock()
            break
        }
        log.Printf("Received: %s\n", message)

        // Broadcast the received message to all clients
        broadcast(message)
    }
}

func main() {
    http.HandleFunc("/ws", wsHandler)
    http.Handle("/", http.FileServer(http.Dir("./public")))
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
