package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/doduykhang/hermes/chat/pkg/dto"
	"github.com/olahol/melody"
)

type WebSocket interface {
	BroadcastToRoom(roomId string, message dto.Message) (error)
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

func (ws *webSocket) BroadcastToRoom(roomId string, message dto.Message) (error) {
	log.Println("broadcasting from", os.Getenv("SERVER"))			
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


