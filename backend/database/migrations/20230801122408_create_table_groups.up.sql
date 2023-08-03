CREATE TABLE groups (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    chat_room BIGINT NOT NULL REFERENCES chat_rooms(id)
);