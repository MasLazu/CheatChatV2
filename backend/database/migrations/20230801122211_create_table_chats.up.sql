CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    sender_email VARCHAR(64) NOT NULL REFERENCES users(email),
    message TEXT NOT NULL,
    chat_room BIGINT NOT NULL REFERENCES chat_rooms(id),
    created_at TIMESTAMP DEFAULT NOW()
);