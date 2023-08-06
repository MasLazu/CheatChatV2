package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MasLazu/CheatChatV2/database"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/model/domain"
)

type ChatsRepository interface {
	Save(ctx context.Context, chat domain.Chat) (int64, error)
	GetPreviewGroupChats(ctx context.Context, userEmail string) ([]model.PreviewGroupChat, error)
	GetPreviewPersonalChats(ctx context.Context, userEmail string) ([]model.PreviewPersonalChat, error)
	GetPersonalChatRoom(ctx context.Context, userEmail1 string, userEmail2 string) (int64, error)
}

type ChatsRepositoryImpl struct {
	databaseConn *sql.DB
}

func NewChatsRepository() *ChatsRepositoryImpl {
	return &ChatsRepositoryImpl{
		databaseConn: database.GetDBConn(),
	}
}

func (repository ChatsRepositoryImpl) Save(ctx context.Context, chat domain.Chat) (int64, error) {
	var id int64
	sql := "INSERT INTO chats (sender_email, message, chat_room, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := repository.databaseConn.QueryRowContext(ctx, sql, chat.SenderEmail, chat.Message, chat.ChatRoom, chat.CreatedAt).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (repository ChatsRepositoryImpl) GetPreviewGroupChats(ctx context.Context, userEmail string) ([]model.PreviewGroupChat, error) {
	var previewChatGroups []model.PreviewGroupChat
	sql := "WITH LastChatPerGroup AS (SELECT g.id AS group_id, g.name AS group_name, c.id AS chat_id, c.sender_email, c.message, c.created_at, ROW_NUMBER() OVER (PARTITION BY g.id ORDER BY c.created_at DESC) AS rn FROM groups g INNER JOIN chat_rooms cr ON g.chat_room = cr.id INNER JOIN chats c ON cr.id = c.chat_room INNER JOIN group_users gu ON g.id = gu.group_id INNER JOIN users u ON gu.user_email = u.email WHERE u.email = $1) SELECT group_id, group_name, chat_id, sender_email, message, created_at FROM LastChatPerGroup WHERE rn = 1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail)
	if err != nil {
		return previewChatGroups, err
	}
	defer row.Close()

	for row.Next() {
		var previewChatGroup model.PreviewGroupChat
		if err := row.Scan(&previewChatGroup.GroupId, &previewChatGroup.GroupName, &previewChatGroup.ChatId, &previewChatGroup.SenderEmail, &previewChatGroup.Message, &previewChatGroup.CreatedAt); err != nil {
			return previewChatGroups, err
		}
		previewChatGroups = append(previewChatGroups, previewChatGroup)
	}
	return previewChatGroups, nil
}

func (repository ChatsRepositoryImpl) GetPreviewPersonalChats(ctx context.Context, userEmail string) ([]model.PreviewPersonalChat, error) {
	var previewPersonalChats []model.PreviewPersonalChat
	sql := "WITH LastChatPerson AS (SELECT CASE WHEN p.user_email_1 = $1 THEN p.user_email_2 ELSE p.user_email_1 END AS email, c.id AS chat_id, c.sender_email, c.message, c.created_at, ROW_NUMBER() OVER (PARTITION BY p.chat_room ORDER BY c.created_at DESC) AS rn FROM personals p INNER JOIN chat_rooms cr ON p.chat_room = cr.id INNER JOIN chats c ON cr.id = c.chat_room INNER JOIN users u ON p.user_email_1 = u.email OR p.user_email_2 = u.email WHERE u.email = $1 ) SELECT email, chat_id, sender_email, message, created_at FROM LastChatPerson WHERE rn = 1"
	row, err := repository.databaseConn.QueryContext(ctx, sql, userEmail)
	if err != nil {
		return previewPersonalChats, err
	}
	defer row.Close()

	for row.Next() {
		var previewPersonalChat model.PreviewPersonalChat
		if err := row.Scan(&previewPersonalChat.Email, &previewPersonalChat.ChatId, &previewPersonalChat.SenderEmail, &previewPersonalChat.Message, &previewPersonalChat.CreatedAt); err != nil {
			return previewPersonalChats, err
		}
		previewPersonalChats = append(previewPersonalChats, previewPersonalChat)
	}
	return previewPersonalChats, nil
}

func (repository ChatsRepositoryImpl) GetPersonalChatRoom(ctx context.Context, userEmail1 string, userEmail2 string) (int64, error) {
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
