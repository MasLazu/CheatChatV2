CREATE TABLE users (
    email VARCHAR(64) PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
