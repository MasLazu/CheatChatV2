package service

import (
	"context"

	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
)

type GroupService interface {
	MakeGroup(userEmail string, groupName string, ctx context.Context) (domain.Group, error)
}

type GroupServiceImpl struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) GroupService {
	return &GroupServiceImpl{
		groupRepository: groupRepository,
	}
}

func (service *GroupServiceImpl) MakeGroup(userEmail string, groupName string, ctx context.Context) (domain.Group, error) {
	var groupEmpty domain.Group
	group, err := service.groupRepository.Save(ctx, groupName)
	if err != nil {
		return groupEmpty, err
	}

	if err := service.groupRepository.AddMemberToGroup(ctx, userEmail, group.Id); err != nil {
		return groupEmpty, err
	}

	return group, nil
}
