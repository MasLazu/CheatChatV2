package repository

import (
	"context"
	"database/sql"

	"github.com/MasLazu/CheatChatV2/database"
	"github.com/MasLazu/CheatChatV2/model/domain"
)

type GroupRepository interface {
	Save(ctx context.Context, name string) (domain.Group, error)
	GetUserGroups(ctx context.Context, userEmail string) ([]domain.Group, error)
	AddMemberToGroup(ctx context.Context, userEmail string, groupId int64) error
	GetUserGroupIds(ctx context.Context, userEmail string) ([]int64, error)
}

type GroupRepositoryImpl struct {
	databaseConn *sql.DB
}

func NewGroupReposiroty() GroupRepository {
	return &GroupRepositoryImpl{
		databaseConn: database.GetDBConn(),
	}
}

func (repository GroupRepositoryImpl) Save(ctx context.Context, name string) (domain.Group, error) {
	var group domain.Group
	sql := "WITH new_chat_room AS (INSERT INTO chat_rooms DEFAULT VALUES RETURNING id) INSERT INTO groups (name, chat_room) VALUES ($1, (SELECT id FROM new_chat_room)) RETURNING id"
	if err := repository.databaseConn.QueryRowContext(ctx, sql, name).Scan(&group.Id); err != nil {
		return group, err
	}

	group.Name = name
	return group, nil
}

func (repository GroupRepositoryImpl) GetUserGroups(ctx context.Context, userEmail string) ([]domain.Group, error) {
	var groups []domain.Group
	sql := "SELECT g.id, g.name FROM users u INNER JOIN group_users gs ON u.email = gs.user_email INNER JOIN groups g ON gs.group_id = g.id WHERE u.email = $1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail)
	if err != nil {
		return groups, err
	}
	defer row.Close()

	for row.Next() {
		group := domain.Group{}
		if err := row.Scan(&group.Id, &group.Name); err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (repository GroupRepositoryImpl) AddMemberToGroup(ctx context.Context, userEmail string, groupId int64) error {
	sql := "INSERT INTO group_users (user_email, group_id) VALUES ($1, $2)"
	if _, err := repository.databaseConn.ExecContext(ctx, sql, userEmail, groupId); err != nil {
		return err
	}

	return nil
}

func (repository GroupRepositoryImpl) GetUserGroupIds(ctx context.Context, userEmail string) ([]int64, error) {
	var groupIds []int64
	sql := "SELECT g.id FROM users u INNER JOIN group_users gs ON u.email = gs.user_email INNER JOIN groups g ON gs.group_id = g.id WHERE u.email = $1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail)
	if err != nil {
		return groupIds, err
	}
	defer row.Close()

	for row.Next() {
		var groupId int64
		if err := row.Scan(&groupId); err != nil {
			return groupIds, err
		}
		groupIds = append(groupIds, groupId)
	}
	return groupIds, nil
}
