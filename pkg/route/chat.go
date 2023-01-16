package route

import (
	"github.com/doduykhang/hermes/chat/pkg/handler"
	"github.com/go-chi/chi/v5"
)

func ChatRoute(r chi.Router, h *handler.Chat) {
	go h.HandleMessage()
	r.Get("/ws/{roomId}", h.HandleConnect)
}
