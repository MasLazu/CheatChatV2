CREATE TABLE contacts (
    name VARCHAR(64) NOT NULL,
    user_email VARCHAR(64) REFERENCES users(email),
    saved_user_email VARCHAR(64) REFERENCES users(email),
    PRIMARY KEY (user_email, saved_user_email)
);