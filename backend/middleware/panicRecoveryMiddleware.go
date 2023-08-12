package middleware

import (
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
)

type PanicRecoveryMiddleware struct{}

func NewPanicRecoveryMiddleware() *PanicRecoveryMiddleware {
	return &PanicRecoveryMiddleware{}
}

func (middleware *PanicRecoveryMiddleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err)
			}
		}()
		next.ServeHTTP(writer, request)
	})
}
