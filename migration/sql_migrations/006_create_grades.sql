-- +migrate StatementBegin
-- +migrate Up
CREATE TABLE grades (
    id SERIAL PRIMARY KEY,
    enrollment_id INTEGER NOT NULL REFERENCES enrollments(id) ON DELETE CASCADE,
    grade VARCHAR(2) CHECK (grade IN ('A', 'B', 'C', 'D', 'E')),
    remarks TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS grades;
-- +migrate StatementEnd
