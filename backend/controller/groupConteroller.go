package controller

import (
	"net/http"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
)

func GetUserGroupsController(writer http.ResponseWriter, request *http.Request) {
	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "something went wrong"})
	}

	groupRepository := repository.NewGroupReposiroty()
	groups, err := groupRepository.GetUserGroups(request.Context(), user.Email)
	if err != nil {
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", model.MessageResponse{Message: "something went wrong"})
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", groups)
}

func MakeGroupController(writer http.ResponseWriter, request *http.Request) {
	groupRequest := model.MakeGroupRequest{}
	if err := helper.ReadRequestBody(request, &groupRequest); err != nil {
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD_REQUEST", model.MessageResponse{Message: "bad request"})
		return
	}

	helper.Validate(writer, groupRequest)

	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "something went wrong"})
	}

	groupService := service.NewGroupService()
	group, err := groupService.MakeGroup(user.Email, groupRequest.Name, request.Context())
	if err != nil {
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", model.MessageResponse{Message: "something went wrong"})
		return
	}

	helper.WriteResponse(writer, http.StatusOK, "OK", group)
}
