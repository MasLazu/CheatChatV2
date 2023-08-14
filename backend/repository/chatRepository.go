package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MasLazu/CheatChatV2/model/web"

	"github.com/MasLazu/CheatChatV2/model/domain"
)

type ChatRepository interface {
	Save(ctx context.Context, chat domain.Chat) (int64, error)
	GetPreviewGroupChats(ctx context.Context, userEmail string) ([]web.PreviewGroupChat, error)
	GetPreviewPersonalChats(ctx context.Context, userEmail string) ([]web.PreviewPersonalChat, error)
	GetPersonalChats(ctx context.Context, userEmail1 string, userEmail2 string) ([]domain.Chat, error)
	GetGroupChats(ctx context.Context, groupId int64) ([]domain.Chat, error)
}

type ChatRepositoryImpl struct {
	databaseConn *sql.DB
}

func NewChatsRepository(databaseConn *sql.DB) ChatRepository {
	return &ChatRepositoryImpl{
		databaseConn: databaseConn,
	}
}

func (repository *ChatRepositoryImpl) Save(ctx context.Context, chat domain.Chat) (int64, error) {
	var id int64
	sql := "INSERT INTO chats (sender_email, message, chat_room, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := repository.databaseConn.QueryRowContext(ctx, sql, chat.SenderEmail, chat.Message, chat.ChatRoom, chat.CreatedAt).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (repository *ChatRepositoryImpl) GetPreviewGroupChats(ctx context.Context, userEmail string) ([]web.PreviewGroupChat, error) {
	var previewChatGroups []web.PreviewGroupChat
	sql := "WITH LastChatPerGroup AS (SELECT g.id AS group_id, g.name AS group_name, c.id AS chat_id, c.sender_email, c.message, c.created_at, ROW_NUMBER() OVER (PARTITION BY g.id ORDER BY c.created_at DESC) AS rn FROM groups g INNER JOIN chat_rooms cr ON g.chat_room = cr.id INNER JOIN chats c ON cr.id = c.chat_room INNER JOIN group_users gu ON g.id = gu.group_id INNER JOIN users u ON gu.user_email = u.email WHERE u.email = $1) SELECT group_id, group_name, chat_id, sender_email, message, created_at FROM LastChatPerGroup WHERE rn = 1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail)
	if err != nil {
		return previewChatGroups, err
	}
	defer row.Close()

	for row.Next() {
		var previewChatGroup web.PreviewGroupChat
		if err := row.Scan(&previewChatGroup.GroupId, &previewChatGroup.GroupName, &previewChatGroup.ChatId, &previewChatGroup.SenderEmail, &previewChatGroup.Message, &previewChatGroup.CreatedAt); err != nil {
			return previewChatGroups, err
		}
		previewChatGroups = append(previewChatGroups, previewChatGroup)
	}
	return previewChatGroups, nil
}

func (repository *ChatRepositoryImpl) GetPreviewPersonalChats(ctx context.Context, userEmail string) ([]web.PreviewPersonalChat, error) {
	var previewPersonalChats []web.PreviewPersonalChat
	sql := "WITH LastChatPerson AS (SELECT CASE WHEN p.user_email_1 = $1 THEN p.user_email_2 ELSE p.user_email_1 END AS email, c.id AS chat_id, c.sender_email, c.message, c.created_at, ROW_NUMBER() OVER (PARTITION BY p.chat_room ORDER BY c.created_at DESC) AS rn FROM personals p INNER JOIN chat_rooms cr ON p.chat_room = cr.id INNER JOIN chats c ON cr.id = c.chat_room INNER JOIN users u ON p.user_email_1 = u.email OR p.user_email_2 = u.email WHERE u.email = $1 ) SELECT email, chat_id, sender_email, message, created_at FROM LastChatPerson WHERE rn = 1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail)
	if err != nil {
		return previewPersonalChats, err
	}
	defer row.Close()

	for row.Next() {
		var previewPersonalChat web.PreviewPersonalChat
		if err := row.Scan(&previewPersonalChat.Email, &previewPersonalChat.ChatId, &previewPersonalChat.SenderEmail, &previewPersonalChat.Message, &previewPersonalChat.CreatedAt); err != nil {
			return previewPersonalChats, err
		}
		previewPersonalChats = append(previewPersonalChats, previewPersonalChat)
	}
	return previewPersonalChats, nil
}

func (repository *ChatRepositoryImpl) GetPersonalChats(ctx context.Context, userEmail1 string, userEmail2 string) ([]domain.Chat, error) {
	var chats []domain.Chat
	sql := "SELECT c.id, c.sender_email, c.message, c.created_at FROM personals p INNER JOIN chat_rooms cr ON p.chat_room = cr.id INNER JOIN chats c ON cr.id = c.chat_room WHERE (p.user_email_1 = $1 AND p.user_email_2 = $2) OR (p.user_email_2 = $1 AND p.user_email_1 = $2);"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail1, userEmail2)
	if err != nil {
		return chats, err
	}
	defer row.Close()

	for row.Next() {
		var chat domain.Chat
		if err := row.Scan(&chat.Id, &chat.SenderEmail, &chat.Message, &chat.CreatedAt); err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}

	if len(chats) == 0 {
		return chats, errors.New("not found")
	}

	return chats, nil
}

func (repository *ChatRepositoryImpl) GetGroupChats(ctx context.Context, groupId int64) ([]domain.Chat, error) {
	var chats []domain.Chat
	sql := "SELECT c.id, c.sender_email, c.message, c.created_at FROM groups g INNER JOIN chat_rooms cr ON cr.id = g.chat_room INNER JOIN chats c ON cr.id = c.chat_room WHERE g.id = $1;"
	row, err := repository.databaseConn.QueryContext(ctx, sql, groupId)
	if err != nil {
		return chats, err
	}
	defer row.Close()

	for row.Next() {
		var chat domain.Chat
		if err := row.Scan(&chat.Id, &chat.SenderEmail, &chat.Message, &chat.CreatedAt); err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}

	if len(chats) == 0 {
		return chats, errors.New("not found")
	}

	return chats, nil
}
