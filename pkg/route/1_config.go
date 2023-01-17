package route

import "github.com/doduykhang/hermes/chat/pkg/middleware"

var (
	authMiddleware = middleware.NewAuthenticateMiddleware()
)
