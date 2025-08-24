//Go-сервер с WebSocket:

package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    "log"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    for {
        // Чтение сообщения от клиента
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }

        // Обработка сообщения и отправка ответа
        log.Printf("Received: %s", message)
        err = conn.WriteMessage(messageType, []byte("Hello from Go server!"))
        if err != nil {
            log.Println("Write error:", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    http.ListenAndServe(":8080", nil)
}