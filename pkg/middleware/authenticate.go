package middleware

import (
	"net/http"

	"github.com/doduykhang/hermes/chat/pkg/service"
)

type AuthenticateMiddleware struct {
	jwtService service.JwtServcie
}

func NewAuthenticateMiddleware () *AuthenticateMiddleware {
	return &AuthenticateMiddleware {
		jwtService: service.NewJwtService(),
	}
}

func (m *AuthenticateMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header["Authorization"]
		if len(authorization) == 0 {
			http.Error(w, "Who are you", http.StatusUnauthorized)
			return
		}
		_, err := m.jwtService.Parse(authorization[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return 
		}

		next.ServeHTTP(w, r)
	})
}
