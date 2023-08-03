package repository

import (
	"context"
	"database/sql"
	"github.com/MasLazu/CheatChatV2/database"
	"github.com/MasLazu/CheatChatV2/model/domain"
)

type ContactRepository interface {
	Save(ctx context.Context, contact domain.Contact) error
	GetAllUserContact(ctx context.Context, userEmail string) ([]domain.Contact, error)
}

type ContactRepositoryImpl struct {
	databaseConn *sql.DB
}

func NewContactRepository() *ContactRepositoryImpl {
	return &ContactRepositoryImpl{
		databaseConn: database.GetDBConn(),
	}
}

func (repository ContactRepositoryImpl) Save(ctx context.Context, contact domain.Contact) error {
	sql := "INSERT INTO contacts(user_email, name, saved_user_email) values ($1, $2, $3)"
	if _, err := repository.databaseConn.ExecContext(ctx, sql, contact.UserEmail, contact.Name, contact.SavedUserEmail); err != nil {
		return err
	}
	return nil
}

func (repository ContactRepositoryImpl) GetUserContacts(ctx context.Context, userEmail string) ([]domain.Contact, error) {
	var contacts []domain.Contact
	sql := "SELECT name, saved_user_email FROM contacts WHERE user_email = $1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail)
	if err != nil {
		return contacts, err
	}
	defer row.Close()

	for row.Next() {
		var contact domain.Contact
		if err := row.Scan(&contact.Name, &contact.SavedUserEmail); err != nil {
			return contacts, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
