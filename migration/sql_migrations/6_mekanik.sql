-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE mekanik (
    id_mekanik     SERIAL PRIMARY KEY,
    nama_mekanik   VARCHAR(255) NOT NULL,
    telp_mekanik   VARCHAR(50),
    keahlian       VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_by VARCHAR(100), 
    modified_at TIMESTAMP, 
    modified_by VARCHAR(100)
);

-- +migrate Down
DROP TABLE IF EXISTS mekanik;

-- +migrate StatementEnd
