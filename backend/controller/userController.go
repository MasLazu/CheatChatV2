package controller

import (
	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/service"
	"net/http"
)

func RegisterController(writer http.ResponseWriter, request *http.Request) {
	userRequest := model.RegisterUserRequest{}
	if err := helper.ReadRequestBody(request, &userRequest); err != nil {
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", model.MessageResponse{Message: "bad request"})
		return
	}

	if err := helper.Validate(writer, userRequest); err != nil {
		return
	}

	userService := service.NewUserService()
	if err := userService.Register(userRequest, request.Context()); err != nil {
		if err.Error() == "something went wrong" {
			helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", model.MessageResponse{Message: err.Error()})
			return
		}
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", model.MessageResponse{Message: err.Error()})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", model.MessageResponse{Message: "register success"})
}

func LoginController(writer http.ResponseWriter, request *http.Request) {
	userRequest := model.LoginUserRequest{}
	if err := helper.ReadRequestBody(request, &userRequest); err != nil {
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", model.MessageResponse{Message: "bad request"})
	}

	if err := helper.Validate(writer, userRequest); err != nil {
		return
	}

	sessionService := service.NewSessionService()
	session, err := sessionService.Login(userRequest, request.Context())

	if err != nil {
		if err.Error() == "something went wrong" {
			helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", model.MessageResponse{Message: err.Error()})
			return
		}
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", model.MessageResponse{Message: err.Error()})
		return
	}

	helper.SetCookies(writer, "session", session.Token, session.ExpiredAt)
	helper.WriteResponse(writer, http.StatusOK, "OK", model.MessageResponse{Message: "login success"})
}

func CurrentController(writer http.ResponseWriter, request *http.Request) {
	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "login only route"})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", model.CuerrentResponse{Email: user.Email, Username: user.Username})
}
