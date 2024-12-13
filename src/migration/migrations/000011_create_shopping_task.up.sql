CREATE TABLE shopping_tasks (
    id SERIAL PRIMARY KEY,
    food_name VARCHAR(255) NOT NULL,
    quantity VARCHAR(255) NOT NULL,
    shopping_list_id INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (shopping_list_id) REFERENCES shopping_lists (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);