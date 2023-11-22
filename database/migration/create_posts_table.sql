CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at timestamptz DEFAULT current_timestamp
);