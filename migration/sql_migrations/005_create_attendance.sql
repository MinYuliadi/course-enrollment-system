-- +migrate StatementBegin
-- +migrate Up
CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    enrollment_id INTEGER NOT NULL REFERENCES enrollments(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status VARCHAR(20) CHECK (status IN ('present', 'absent', 'late'))
);
-- +migrate Down
DROP TABLE IF EXISTS attendance;
-- +migrate StatementEnd
