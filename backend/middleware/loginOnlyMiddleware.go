package middleware

import (
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model/web"
	"github.com/MasLazu/CheatChatV2/service"
)

type LoginOnlyMiddleware struct {
	sessionService *service.SessionService
}

func NewLoginOnlyMiddleware(sessionService *service.SessionService) *LoginOnlyMiddleware {
	return &LoginOnlyMiddleware{
		sessionService: sessionService,
	}
}

func (middleware *LoginOnlyMiddleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		_, err := middleware.sessionService.Current(request, request.Context())
		if err == nil {
			next.ServeHTTP(writer, request)
			return
		}

		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", web.MessageResponse{Message: "login oly route"})
	})
}
