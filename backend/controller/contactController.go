package controller

import (
	"log"
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
)

func AddContactController(writer http.ResponseWriter, request *http.Request) {
	contactRequest := model.AddContactRequest{}
	if err := helper.ReadRequestBody(request, &contactRequest); err != nil {
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD_REQUEST", model.MessageResponse{Message: "bad request"})
		return
	}

	helper.Validate(writer, contactRequest)

	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "something went wrong"})
		return
	}

	contact := domain.Contact{
		UserEmail:      user.Email,
		Name:           contactRequest.Name,
		SavedUserEmail: contactRequest.Email,
	}

	contactService := service.NewContactService()
	if err := contactService.AddContact(contact, request.Context()); err != nil {
		log.Println(err)
		if err.Error() == "user not found" {
			helper.WriteResponse(writer, http.StatusNotFound, "NOT_FOUND", model.MessageResponse{Message: "user not found"})
			return
		} else if err.Error() == "pq: duplicate key value violates unique constraint \"contacts_pkey\"" {
			helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", model.MessageResponse{Message: "contact already exist"})
			return
		}
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", model.MessageResponse{Message: err.Error()})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", model.AddContactRequest{Name: contact.Name, Email: contact.SavedUserEmail})
}

func GetContactUserController(writer http.ResponseWriter, request *http.Request) {
	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "something went wrong"})
		return
	}

	contactRepository := repository.NewContactRepository()
	contacts, err := contactRepository.GetUserContacts(request.Context(), user.Email)
	if err != nil {
		log.Println(err)
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", model.MessageResponse{Message: "something went wrong"})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", contacts)
}
