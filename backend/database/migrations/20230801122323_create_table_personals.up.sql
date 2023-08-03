CREATE TABLE personals (
    user_email_1 VARCHAR(64) REFERENCES users(email),
    user_email_2 VARCHAR(64) REFERENCES users(email),
    chat_room BIGINT NOT NULL REFERENCES chat_rooms(id),
    PRIMARY KEY (user_email_1, user_email_2)
);