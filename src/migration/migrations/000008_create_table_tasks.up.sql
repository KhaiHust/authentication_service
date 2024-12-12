CREATE TABLE shopping_lists (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    DESCRIPTION TEXT,
    created_by INT NOT NULL,
    assigned_to INT NOT NULL,
    PRIMARY KEY (created_by) REFERENCES users (id),
    PRIMARY KEY (assigned_to) REFERENCES users (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shopping_list_groups (
    id SERIAL PRIMARY KEY,
    shopping_list_id INT NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (shopping_list_id) REFERENCES shopping_lists (id),
    FOREIGN KEY (group_id) REFERENCES groups (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

