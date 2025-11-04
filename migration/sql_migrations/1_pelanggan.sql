-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE pelanggan (
    id     SERIAL PRIMARY KEY,
    nama_pelanggan   VARCHAR(255) NOT NULL,
    alamat           VARCHAR(255),
    telp_pelanggan   VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_by VARCHAR(100), 
    modified_at TIMESTAMP, 
    modified_by VARCHAR(100)
);

-- +migrate Down
DROP TABLE IF EXISTS pelanggan;

-- +migrate StatementEnd
