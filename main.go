package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// WebSocket bağlantısı yükseltme
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Sonsuz döngü ile gelen verileri oku ve geri gönder
	for {
		// Veri okuma
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Veri yazma
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	fmt.Println("Websocket sunucusu 8080 portunda çalışıyor...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
