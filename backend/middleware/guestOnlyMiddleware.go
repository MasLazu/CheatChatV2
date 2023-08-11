package middleware

import (
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/service"
)

type GuestOnlyMiddleware struct {
	sessionService service.SessionService
}

func NewGuestOnlyMiddleware(sessionService service.SessionService) *GuestOnlyMiddleware {
	return &GuestOnlyMiddleware{
		sessionService: sessionService,
	}
}

func (middleware GuestOnlyMiddleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if _, err := middleware.sessionService.Current(request, request.Context()); err != nil {
			next.ServeHTTP(writer, request)
			return
		}

		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "guest only route"})
	})
}
