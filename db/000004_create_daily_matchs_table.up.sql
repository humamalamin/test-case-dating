CREATE TABLE daily_matches
(
    id VARCHAR(64) NOT NULL,
    first_user_id VARCHAR(64) NOT NULL,
    second_user_id VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (first_user_id) REFERENCES users(id),
    FOREIGN KEY (second_user_id) REFERENCES users(id)
);