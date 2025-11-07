-- +migrate StatementBegin
-- +migrate Up
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    teacher_id INTEGER REFERENCES teachers(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS courses;
-- +migrate StatementEnd
