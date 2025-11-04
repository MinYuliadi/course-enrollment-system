-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE kendaraan (
    id          SERIAL PRIMARY KEY,
    id_pelanggan      INTEGER NOT NULL,
    plat_kendaraan        VARCHAR(50),
    tahun_produksi    INTEGER,
    status_kepemilikan VARCHAR(100),
    merek             VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_by VARCHAR(100), 
    modified_at TIMESTAMP, 
    modified_by VARCHAR(100),
    CONSTRAINT fk_kendaraan_pelanggan FOREIGN KEY (id_pelanggan)
        REFERENCES pelanggan (id)
        ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS kendaraan;

-- +migrate StatementEnd
