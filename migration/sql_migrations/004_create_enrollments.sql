-- +migrate StatementBegin
-- +migrate Up
CREATE TABLE enrollments (
    id SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (student_id, course_id)
);
-- +migrate Down
DROP TABLE IF EXISTS enrollments;
-- +migrate StatementEnd
