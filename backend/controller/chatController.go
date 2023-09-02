package controller

import (
	"net/http"
	"strconv"

	"github.com/MasLazu/CheatChatV2/model/web"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
	"github.com/gorilla/mux"
)

type ChatController struct {
	sessionService *service.SessionService
	chatRepository *repository.ChatRepository
}

func NewChatController(sessionService *service.SessionService, chatRepository *repository.ChatRepository) *ChatController {
	return &ChatController{
		sessionService: sessionService,
		chatRepository: chatRepository,
	}
}

func (controller *ChatController) GetPreviews(writer http.ResponseWriter, request *http.Request) {
	user, err := controller.sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteUnauthorizedError(writer)
		return
	}

	previewGroupChat, err := controller.chatRepository.GetPreviewGroupChats(request.Context(), user.Email)
	if err != nil {
		helper.WriteInternalServerError(writer)
		return
	}

	previewPersonalChat, err := controller.chatRepository.GetPreviewPersonalChats(request.Context(), user.Email)
	if err != nil {
		helper.WriteInternalServerError(writer)
		return
	}

	helper.WriteOk(writer, web.PreviewChatResponse{Group: previewGroupChat, Personal: previewPersonalChat})
}

func (controller *ChatController) GetPersonals(writer http.ResponseWriter, request *http.Request) {
	email := mux.Vars(request)["email"]

	user, err := controller.sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteUnauthorizedError(writer)
		return
	}

	chats, err := controller.chatRepository.GetPersonalChats(request.Context(), user.Email, email)
	if err != nil {
		helper.WriteNotFoundError(writer)
		return
	}

	helper.WriteOk(writer, chats)
}

func (controller *ChatController) GetGroups(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	if err != nil {
		helper.WriteNotFoundError(writer)
		return
	}

	chats, err := controller.chatRepository.GetGroupChats(request.Context(), id)
	if err != nil {
		helper.WriteNotFoundError(writer)
		return
	}

	helper.WriteOk(writer, chats)
}
