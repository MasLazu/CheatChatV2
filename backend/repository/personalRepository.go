package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MasLazu/CheatChatV2/model/domain"
)

type PersonalRepository interface {
	Save(userEmail1 string, userEmail2 string, ctx context.Context) error
	GetByMember(userEmail1 string, userEmail2 string, ctx context.Context) (domain.Personal, error)
}

type PersonalRepositoryImpl struct {
	databaseConn *sql.DB
}

func NewPersonalRepository(databaseConn *sql.DB) PersonalRepository {
	return &PersonalRepositoryImpl{
		databaseConn: databaseConn,
	}
}

func (repository *PersonalRepositoryImpl) Save(userEmail1 string, userEmail2 string, ctx context.Context) error {
	sql := "WITH new_chat_room AS (INSERT INTO chat_rooms DEFAULT VALUES RETURNING id) INSERT INTO personals (user_email_1, user_email_2, chat_room) VALUES ($1, $2, (SELECT id FROM new_chat_room))"
	if _, err := repository.databaseConn.ExecContext(ctx, sql, userEmail1, userEmail2); err != nil {
		return err
	}
	return nil
}

func (repository *PersonalRepositoryImpl) GetByMember(userEmail1 string, userEmail2 string, ctx context.Context) (domain.Personal, error) {
	sql := "SELECT user_email_1, user_email_2, chat_room FROM personals WHERE (user_email_1 = $1 AND user_email_2 = $2) OR (user_email_1 = $2 AND user_email_2 = $1)"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail1, userEmail2)
	personal := domain.Personal{}
	if err != nil {
		return personal, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&personal.UserEmail1, &personal.UserEmail2, &personal.ChatRoom); err != nil {
			return personal, err
		}
		return personal, nil
	}
	return personal, errors.New("not found")
}
