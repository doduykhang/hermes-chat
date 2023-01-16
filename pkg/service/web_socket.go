package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olahol/melody"
)

type Message struct {
	Message string `json:"message"`
}

type WebSocket interface {
	BroadcastToRoom(roomId string, message Message) (error)
	Handle(w http.ResponseWriter, r *http.Request) (error)
}

type webSocket struct {
	melody *melody.Melody
}

func NewWebSocket (melody *melody.Melody) WebSocket {
	return &webSocket{
		melody: melody,	
	}	
}

func (ws *webSocket) BroadcastToRoom(roomId string, message Message) (error) {
	messageByte, err := json.Marshal(&message)
	if err != nil {
		return err
	}

	err = ws.melody.BroadcastFilter(messageByte, func(q *melody.Session) bool {
		return q.Request.URL.Path == "/ws/" +roomId
	})

	if err != nil {
		fmt.Printf("broadcast err, %v", err)
	}
	return err
}

func (ws *webSocket) Handle(w http.ResponseWriter, r *http.Request) (error) {
	return ws.melody.HandleRequest(w,r )
}


