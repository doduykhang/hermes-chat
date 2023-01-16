package dto

import "time"

type IncomingMessage struct {
	RoomId string `json:"roomId"`
	Message Message `json:"message"`
}

type Message struct {
	Message string `json:"message"`
	Avatar string `json:"avatar"`
	Username string `json:"userName"`
	Timestamp time.Time `json:"timestamp"`
}
