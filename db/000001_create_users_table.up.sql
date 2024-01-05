CREATE TABLE users
(
    id VARCHAR(64) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NULL,
    email VARCHAR(100) NOT NULL,
    gender INT(1) NOT NULL,
    password VARCHAR(200) NOT NULL UNIQUE,
    birth_date DATETIME NULL,
    location VARCHAR(100) NULL,
    lat FLOAT(20, 0) NULL DEFAULT 0,
    lng FLOAT(20, 0) NULL DEFAULT 0,
    is_premium TINYINT(1) NULL DEFAULT 0,
    verified_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);