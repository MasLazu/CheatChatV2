package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MasLazu/CheatChatV2/database"
	"github.com/MasLazu/CheatChatV2/model/domain"
)

type SessionRepository interface {
	InsertIfExistUpdate(session domain.Session, ctx context.Context) error
	GetByToken(ctx context.Context, token string) (domain.Session, error)
}

type SessionRepositoryImpl struct {
	databaseConn *sql.DB
}

func NewSessionRepository() SessionRepository {
	return &SessionRepositoryImpl{
		databaseConn: database.GetDBConn(),
	}
}

func (repository SessionRepositoryImpl) InsertIfExistUpdate(session domain.Session, ctx context.Context) error {
	sql := "INSERT INTO sessions (user_email, token, expired_at) VALUES ($1, $2, $3) ON CONFLICT (user_email) DO UPDATE SET user_email = $1, token = $2, expired_at = $3"
	if _, err := repository.databaseConn.ExecContext(ctx, sql, session.UserEmail, session.Token, session.ExpiredAt); err != nil {
		return err
	}
	return nil
}

func (repository SessionRepositoryImpl) GetByToken(ctx context.Context, token string) (domain.Session, error) {
	sql := "SELECT user_email, token, expired_at FROM sessions WHERE token = $1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, token)
	session := domain.Session{}
	if err != nil {
		return session, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&session.UserEmail, &session.Token, &session.ExpiredAt); err != nil {
			return session, err
		}
		return session, nil
	}
	return session, errors.New("not found")
}
