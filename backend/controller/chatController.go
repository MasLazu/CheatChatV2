package controller

import (
	"github.com/MasLazu/CheatChatV2/model/web"
	"net/http"
	"strconv"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
	"github.com/gorilla/mux"
)

type ChatController interface {
	GetPreviews(writer http.ResponseWriter, request *http.Request)
	GetPersonals(writer http.ResponseWriter, request *http.Request)
	GetGroups(writer http.ResponseWriter, request *http.Request)
}

type ChatControllerImpl struct {
	sessionService service.SessionService
	chatRepository repository.ChatRepository
}

func NewChatController(sessionService service.SessionService, chatRepository repository.ChatRepository) ChatController {
	return &ChatControllerImpl{
		sessionService: sessionService,
		chatRepository: chatRepository,
	}
}

func (controller *ChatControllerImpl) GetPreviews(writer http.ResponseWriter, request *http.Request) {
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

func (controller *ChatControllerImpl) GetPersonals(writer http.ResponseWriter, request *http.Request) {
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

func (controller *ChatControllerImpl) GetGroups(writer http.ResponseWriter, request *http.Request) {
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
