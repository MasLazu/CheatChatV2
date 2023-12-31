package controller

import (
	"net/http"

	"github.com/MasLazu/CheatChatV2/model/web"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
)

type GroupController struct {
	sessionService  *service.SessionService
	groupService    *service.GroupService
	groupRepository *repository.GroupRepository
}

func NewGroupController(sessionService *service.SessionService, groupService *service.GroupService, groupRepository *repository.GroupRepository) *GroupController {
	return &GroupController{
		sessionService:  sessionService,
		groupService:    groupService,
		groupRepository: groupRepository,
	}
}

func (controller *GroupController) GetUserGroups(writer http.ResponseWriter, request *http.Request) {
	user, err := controller.sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteUnauthorizedError(writer)
		return
	}

	groups, err := controller.groupRepository.GetUserGroups(request.Context(), user.Email)
	if err != nil {
		helper.WriteInternalServerError(writer)
		return
	}

	helper.WriteOk(writer, groups)
}

func (controller *GroupController) Make(writer http.ResponseWriter, request *http.Request) {
	groupRequest := web.MakeGroupRequest{}
	if err := helper.ReadRequestBody(request, &groupRequest); err != nil {
		helper.WriteBadRequestError(writer)
		return
	}

	if err := helper.Validate(writer, groupRequest); err != nil {
		return
	}

	user, err := controller.sessionService.Current(request, request.Context())
	if err != nil {
		helper.WriteUnauthorizedError(writer)
		return
	}

	group, err := controller.groupService.MakeGroup(user.Email, groupRequest.Name, request.Context())
	if err != nil {
		helper.WriteInternalServerError(writer)
		return
	}

	helper.WriteOk(writer, group)
}
