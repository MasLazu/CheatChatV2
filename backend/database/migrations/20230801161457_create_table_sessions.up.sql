CREATE TABLE sessions (
    user_email VARCHAR(64) PRIMARY KEY REFERENCES users(email),
    token VARCHAR(64) NOT NULL UNIQUE,
    expired_at TIMESTAMP DEFAULT (NOW() + INTERVAL '1 DAY')
);