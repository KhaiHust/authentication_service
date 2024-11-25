CREATE TABLE refresh_tokens (
                                id SERIAL PRIMARY KEY,
                                user_id BIGINT NOT NULL,
                                refresh_token VARCHAR(500) NOT NULL,
                                ip_address VARCHAR(45),
                                user_agent TEXT,
                                expired_at TIMESTAMP NOT NULL,
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);