package httpserver

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type playerServerWS struct {
	conn *websocket.Conn
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) playerServerWS {

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("problem upgrading connection to WebSocket, %v", err)
	}

	return playerServerWS{conn}

}

func (s *playerServerWS) waitForMessage() string {
	_, msg, err := s.conn.ReadMessage()
	if err != nil {
		log.Printf("problem reading message, %v", err)
	}
	return string(msg)
}
