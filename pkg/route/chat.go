package route

import (
	"net/http"

	"github.com/doduykhang/hermes/chat/pkg/handler"
	"github.com/go-chi/chi/v5"
)

func ChatRoute(r chi.Router, h *handler.Chat) {
	go h.HandleMessage()
	r.With(authMiddleware.Authenticate).Get("/ws/{roomId}", h.HandleConnect)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
}
