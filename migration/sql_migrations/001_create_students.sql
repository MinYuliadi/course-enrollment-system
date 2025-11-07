-- +migrate StatementBegin
-- +migrate Up
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE IF EXISTS students;
-- +migrate StatementEnd
