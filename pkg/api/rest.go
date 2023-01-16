package api

import (
	"log"
	"net/http"
	"os"

	"github.com/doduykhang/hermes/chat/pkg/config"
	"github.com/doduykhang/hermes/chat/pkg/handler"
	"github.com/doduykhang/hermes/chat/pkg/route"
	"github.com/doduykhang/hermes/chat/pkg/service"
	"github.com/go-chi/chi/v5"
)

func Serve() {
	conf := config.GetEnv()
	//router and port
	port := conf.Port
	if customPort := os.Getenv("PORT"); customPort != "" {
		port = customPort
	} 
	r := chi.NewRouter()		

	//framework
	redis := config.NewRedisClient(conf.Redis)
	melody := config.NewMelody()

	//services
	ws := service.NewWebSocket(melody)
	subService := service.NewSub(redis)

	//handlers
	chatHandler := handler.NewChat(ws, subService)	
		
	//routes
	route.ChatRoute(r, chatHandler)

	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
