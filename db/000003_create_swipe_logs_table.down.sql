CREATE TABLE swipe_logs
(
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    profile_id VARCHAR(64) NOT NULL,
    swipe_type VARCHAR(20) NOT NULL DEFAULT 'PASS',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (profile_id) REFERENCES profiles(id)
);