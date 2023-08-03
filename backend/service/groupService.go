package service

import (
	"context"
	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
)

type GroupService interface {
}

type GroupServiceImpl struct {
}

func NewGroupService() *GroupServiceImpl {
	return &GroupServiceImpl{}
}

func (GroupServiceImpl) MakeGroup(userEmail string, groupName string, ctx context.Context) (domain.Group, error) {
	var groupEmpty domain.Group
	groupRepository := repository.NewGroupReposiroty()
	group, err := groupRepository.Save(ctx, groupName)
	if err != nil {
		return groupEmpty, err
	}
	if err := groupRepository.AddMemberToGroup(ctx, userEmail, group.Id); err != nil {
		return groupEmpty, err
	}

	return group, nil
}
