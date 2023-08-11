package middleware

import (
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/service"
)

func LoginOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		sessionService := service.NewSessionService()

		_, err := sessionService.Current(request, request.Context())
		if err == nil {
			next.ServeHTTP(writer, request)
			return
		}

		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "login oly route"})
	})
}
