CREATE TABLE premiums
(
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    purchase_type INTEGER(2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);