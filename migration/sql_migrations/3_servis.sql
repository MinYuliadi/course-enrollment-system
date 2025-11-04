-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE servis (
    id          SERIAL PRIMARY KEY,
    tgl_servis         DATE,
    keluhan            TEXT,
    status_pengerjaan  VARCHAR(100),
    id_kendaraan           INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_by VARCHAR(100), 
    modified_at TIMESTAMP, 
    modified_by VARCHAR(100),
    CONSTRAINT fk_servis_kendaraan FOREIGN KEY (id_kendaraan)
        REFERENCES kendaraan (id)
        ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS servis;

-- +migrate StatementEnd
