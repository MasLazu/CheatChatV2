package controller

import (
	"github.com/MasLazu/CheatChatV2/model/web"
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/service"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
	Current(writer http.ResponseWriter, request *http.Request)
}

type UserControllerImpl struct {
	sessionService service.SessionService
	userService    service.UserService
}

func NewUserController(sessionService service.SessionService, userService service.UserService) UserController {
	return &UserControllerImpl{
		sessionService: sessionService,
		userService:    userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	userRequest := web.RegisterUserRequest{}
	if err := helper.ReadRequestBody(request, &userRequest); err != nil {
		helper.WriteBadRequestError(writer)
		return
	}

	if err := helper.Validate(writer, userRequest); err != nil {
		return
	}

	if err := controller.userService.Register(userRequest, request.Context()); err != nil {
		if err.Error() == "something went wrong" {
			helper.WriteInternalServerError(writer)
			return
		}
		helper.WriteBadRequestError(writer)
		return
	}

	helper.WriteOk(writer, web.MessageResponse{Message: "register success"})
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	userRequest := web.LoginUserRequest{}
	if err := helper.ReadRequestBody(request, &userRequest); err != nil {
		helper.WriteBadRequestError(writer)
		return
	}

	if err := helper.Validate(writer, userRequest); err != nil {
		return
	}

	session, err := controller.sessionService.Login(userRequest, request.Context())
	if err != nil {
		if err.Error() == "something went wrong" {
			helper.WriteInternalServerError(writer)
			return
		}
		helper.WriteBadRequestError(writer)
		return
	}

	helper.SetCookies(writer, "session", session.Token, session.ExpiredAt)
	helper.WriteOk(writer, web.MessageResponse{Message: "login success"})
}

func (controller *UserControllerImpl) Current(writer http.ResponseWriter, request *http.Request) {
	user, err := controller.sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteUnauthorizedError(writer)
		return
	}

	helper.WriteOk(writer, web.CuerrentResponse{Email: user.Email, Username: user.Username})
}
