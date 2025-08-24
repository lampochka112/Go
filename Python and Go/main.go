package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
    "bytes"
)

type Request struct {
    Name string `json:"name"`
}

type Response struct {
    Message string `json:"message"`
    Length  int    `json:"length"`
}

func main() {
    // Данные для отправки в Python-скрипт
    req := Request{Name: "World"}
    reqBytes, _ := json.Marshal(req)

    // Запускаем Python-интерпретатор со скриптом
    cmd := exec.Command("python3", "main.py")
    // Создаем буфер для stdin и stdout
    var stdout, stderr bytes.Buffer
    cmd.Stdin = bytes.NewReader(reqBytes)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    // Выполняем команду
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Error running script: %v\nStderr: %s", err, stderr.String())
    }

    // Парсим ответ от Python-скрипта
    var resp Response
    json.Unmarshal(stdout.Bytes(), &resp)

    fmt.Printf("Message: %s\n", resp.Message)
    fmt.Printf("Length: %d\n", resp.Length)
}