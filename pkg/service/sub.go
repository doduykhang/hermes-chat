package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/doduykhang/hermes/chat/pkg/dto"
	"github.com/go-redis/redis/v8"
)

type Sub interface {
	Subscribe(out chan dto.IncomingMessage) 
}

type redisSub struct {
	redis *redis.Client 
}

func NewSub (redis *redis.Client) Sub {
	return &redisSub {
		redis: redis,
	}
}

func (s *redisSub) Subscribe(out chan dto.IncomingMessage) {
	subscriber := s.redis.Subscribe(context.TODO(), "messages")		

	for {
        	msg, err := subscriber.ReceiveMessage(context.TODO())
        	if err != nil {
			fmt.Println(err)
        	}

		var incomingMessage dto.IncomingMessage
		err = json.Unmarshal([]byte(msg.Payload), &incomingMessage)
		
        	if err != nil {
			fmt.Println(err)
        	}
		
		out <- incomingMessage
    	}
}


