package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MasLazu/CheatChatV2/model/domain"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}

type UserRepositoryImpl struct {
	databaseConn *sql.DB
}

func NewUserRepository(databaseConn *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		databaseConn: databaseConn,
	}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, user domain.User) error {
	sql := "INSERT INTO users(email, username, password, created_at) values ($1,$2,$3,NOW())"
	if _, err := repository.databaseConn.ExecContext(ctx, sql, user.Email, user.Username, user.Password); err != nil {
		return err
	}
	return nil
}

func (repository *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	sql := "SELECT email, username, password, created_at FROM users WHERE email=$1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, email)
	user := domain.User{}
	if err != nil {
		return user, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&user.Email, &user.Username, &user.Password, &user.CreatedAt); err != nil {
			return user, err
		}
		return user, nil
	}
	return user, errors.New("not found")
}
