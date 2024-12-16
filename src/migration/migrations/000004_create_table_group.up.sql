CREATE TABLE groups (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        description TEXT,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE group_roles (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR(255) NOT NULL,
                             code VARCHAR(255) NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE group_members (
                              id SERIAL PRIMARY KEY,
                              group_id BIGINT NOT NULL,
                              user_id BIGINT NOT NULL,
                              role_id BIGINT NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
                              FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                              FOREIGN KEY (role_id) REFERENCES group_roles(id) ON DELETE CASCADE
);

