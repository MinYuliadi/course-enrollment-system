-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE sparepart (
    id_sparepart    SERIAL PRIMARY KEY,
    nama_sparepart  VARCHAR(255) NOT NULL,
    harga_satuan    INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_by VARCHAR(100), 
    modified_at TIMESTAMP, 
    modified_by VARCHAR(100)
);

-- +migrate Down
DROP TABLE IF EXISTS sparepart;

-- +migrate StatementEnd
