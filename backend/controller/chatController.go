package controller

import (
	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetPreviewChatController(writer http.ResponseWriter, request *http.Request) {
	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "login only route"})
		return
	}

	chatRepository := repository.NewChatsRepository()
	previewGroupChat, err := chatRepository.GetPreviewGroupChats(request.Context(), user.Email)
	if err != nil {
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", model.MessageResponse{Message: "something went wrong"})
		return
	}
	previewPersonalChat, err := chatRepository.GetPreviewPersonalChats(request.Context(), user.Email)
	if err != nil {
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", model.MessageResponse{Message: "something went wrong"})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", model.PreviewChatResponse{Group: previewGroupChat, Personal: previewPersonalChat})
}

func GetPersonalFullChat(writer http.ResponseWriter, request *http.Request) {
	email := mux.Vars(request)["email"]

	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "login only route"})
		return
	}

	chatRepository := repository.NewChatsRepository()
	chats, err := chatRepository.GetPersonalChats(request.Context(), user.Email, email)
	if err != nil {
		helper.WriteResponse(writer, http.StatusNotFound, "NOT FOUND", model.MessageResponse{Message: "not found"})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", chats)
}

func GetGroupFullChats(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	if err != nil {
		helper.WriteResponse(writer, http.StatusNotFound, "NOT FOUND", model.MessageResponse{Message: "not found"})
		return
	}

	chatRepository := repository.NewChatsRepository()
	chats, err := chatRepository.GetGroupChats(request.Context(), id)
	if err != nil {
		helper.WriteResponse(writer, http.StatusNotFound, "NOT FOUND", model.MessageResponse{Message: "not found"})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", chats)
}
