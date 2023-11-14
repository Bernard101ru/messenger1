// server/server.go
package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var templates = template.Must(template.ParseFiles("templates/template.html"))

// HandleWebPage обработчик для веб-интерфейса
func HandleWebPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "template.html", nil)
}

// HandleConnections обработчик WebSocket соединения
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Бесконечный цикл для чтения сообщений от клиента
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// Отправляем полученное сообщение всем подключенным клиентам
		if err := conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}
