CREATE TABLE profiles
(
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    description TEXT NULL,
    image LONGTEXT NULL,
    interest LONGTEXT NULL,
    preference_match TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);