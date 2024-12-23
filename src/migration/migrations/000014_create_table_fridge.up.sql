CREATE TABLE fridge_items (
    id SERIAL PRIMARY KEY,
    expired_date DATE NOT NULL,
    quantity INT NOT NULL,
    note TEXT,
    created_by INT NOT NULL,
    food_id INT NOT NULL,
    FOREIGN KEY (created_by) REFERENCES users(id),
    FOREIGN KEY (food_id) REFERENCES foods(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)