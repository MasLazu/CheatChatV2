package controller

import (
	"github.com/MasLazu/CheatChatV2/model/web"
	"log"
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
)

type ContactController interface {
	Add(writer http.ResponseWriter, request *http.Request)
	GetUserContacts(writer http.ResponseWriter, request *http.Request)
}

type ContactControllerImpl struct {
	sessionService    service.SessionService
	contactService    service.ContactService
	contactRepository repository.ContactRepository
}

func NewContactController(sessionService service.SessionService, contactService service.ContactService, contactRepository repository.ContactRepository) ContactController {
	return &ContactControllerImpl{
		sessionService:    sessionService,
		contactService:    contactService,
		contactRepository: contactRepository,
	}
}

func (controller *ContactControllerImpl) Add(writer http.ResponseWriter, request *http.Request) {
	contactRequest := web.AddContactRequest{}
	if err := helper.ReadRequestBody(request, &contactRequest); err != nil {
		helper.WriteBadRequestError(writer)
		return
	}

	if err := helper.Validate(writer, contactRequest); err != nil {
		return
	}

	user, err := controller.sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteInternalServerError(writer)
		return
	}

	contact := domain.Contact{
		UserEmail:      user.Email,
		Name:           contactRequest.Name,
		SavedUserEmail: contactRequest.Email,
	}

	if err := controller.contactService.AddContact(contact, request.Context()); err != nil {
		log.Println(err)
		if err.Error() == "user not found" {
			helper.WriteResponse(writer, http.StatusNotFound, "NOT_FOUND", web.MessageResponse{Message: "user not found"})
			return
		} else if err.Error() == "pq: duplicate key value violates unique constraint \"contacts_pkey\"" {
			helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", web.MessageResponse{Message: "contact already exist"})
			return
		}
		helper.WriteInternalServerError(writer)
		return
	}

	helper.WriteOk(writer, web.AddContactRequest{Name: contact.Name, Email: contact.SavedUserEmail})
}

func (controller *ContactControllerImpl) GetUserContacts(writer http.ResponseWriter, request *http.Request) {
	user, err := controller.sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteUnauthorizedError(writer)
		return
	}

	contacts, err := controller.contactRepository.GetUserContacts(request.Context(), user.Email)
	if err != nil {
		helper.WriteInternalServerError(writer)
		return
	}

	helper.WriteOk(writer, contacts)
}
