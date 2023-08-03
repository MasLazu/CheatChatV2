CREATE TABLE group_users (
    user_email VARCHAR(64) REFERENCES users(email),
    group_id BIGINT REFERENCES groups(id),
    PRIMARY KEY (user_email, group_id)
);