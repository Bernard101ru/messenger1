// main.go
package main

import (
	"awesomeProject/server"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Chat Server...")

	// Установка обработчика для веб-интерфейса
	http.HandleFunc("/", server.HandleWebPage)
	// Установка обработчика для WebSocket соединений
	http.HandleFunc("/ws", server.HandleConnections)

	// Запуск сервера на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
