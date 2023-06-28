CREATE TABLE IF NOT EXISTS users_users (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    email VARCHAR(255),
    phone_prefix VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    country_code VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    last_name VARCHAR(255),
    is_confirmed TINYINT(1),

    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,

    PRIMARY KEY(id)
);
