package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type IncomingMessage struct {
	RoomId string `json:"roomId"`
	Message string `json:"message"`
}

type Sub interface {
	Subscribe(out chan IncomingMessage) 
}

type redisSub struct {
	redis *redis.Client 
}

func NewSub (redis *redis.Client) Sub {
	return &redisSub {
		redis: redis,
	}
}

func (s *redisSub) Subscribe(out chan IncomingMessage) {
	subscriber := s.redis.Subscribe(context.TODO(), "messages")		

	for {
        	msg, err := subscriber.ReceiveMessage(context.TODO())
        	if err != nil {
			fmt.Println(err)
        	}

		var incomingMessage IncomingMessage
		err = json.Unmarshal([]byte(msg.Payload), &incomingMessage)
		
        	if err != nil {
			fmt.Println(err)
        	}
		
		out <- incomingMessage
    	}
}


