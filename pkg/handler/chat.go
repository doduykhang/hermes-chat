package handler

import (
	"net/http"

	"github.com/doduykhang/hermes/chat/pkg/service"
)

type Chat struct {
	socket service.WebSocket	
	sub service.Sub
}


func (c *Chat) HandleConnect(w http.ResponseWriter, r *http.Request) {
	c.socket.Handle(w, r)
}

func (c *Chat) HandleMessage() {
	messages := make(chan service.IncomingMessage)
	go c.sub.Subscribe(messages)	
	for message := range messages {
		c.socket.BroadcastToRoom(message.RoomId, service.Message{ Message: message.Message })
	}
}

func NewChat(socket service.WebSocket, sub service.Sub) *Chat {	
	return &Chat {
		socket: socket,
		sub: sub,
	}
}
