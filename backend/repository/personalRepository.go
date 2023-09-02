package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MasLazu/CheatChatV2/model/domain"
)

type PersonalRepository struct {
	databaseConn *sql.DB
}

func NewPersonalRepository(databaseConn *sql.DB) *PersonalRepository {
	return &PersonalRepository{
		databaseConn: databaseConn,
	}
}

func (repository *PersonalRepository) Save(userEmail1 string, userEmail2 string, ctx context.Context) error {
	sql := "WITH new_chat_room AS (INSERT INTO chat_rooms DEFAULT VALUES RETURNING id) INSERT INTO personals (user_email_1, user_email_2, chat_room) VALUES ($1, $2, (SELECT id FROM new_chat_room))"
	if _, err := repository.databaseConn.ExecContext(ctx, sql, userEmail1, userEmail2); err != nil {
		return err
	}
	return nil
}

func (repository *PersonalRepository) GetByMember(userEmail1 string, userEmail2 string, ctx context.Context) (domain.Personal, error) {
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

func (repository *PersonalRepository) GetChatRoom(ctx context.Context, userEmail1 string, userEmail2 string) (int64, error) {
	var chatRoom int64
	sql := "SELECT cr.id from personals p INNER JOIN chat_rooms cr on p.chat_room = cr.id WHERE (p.user_email_1 = $1 AND p.user_email_2 =$2) OR (p.user_email_2 = $1 AND p.user_email_1 = $2)"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail1, userEmail2)
	if err != nil {
		return chatRoom, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&chatRoom); err != nil {
			return chatRoom, err
		}
		return chatRoom, nil
	}
	return chatRoom, errors.New("not found")
}
