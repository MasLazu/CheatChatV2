package middleware

import (
	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/service"
	"net/http"
)

func LoginOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		sessionService := service.NewSessionService()

		if _, err := sessionService.Current(request, request.Context()); err == nil {
			next.ServeHTTP(writer, request)
			return
		}

		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "login oly route"})
	})
}
