-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE detail (
    id_pengerjaan      SERIAL PRIMARY KEY,
    id_servis          INTEGER NOT NULL,
    id_sparepart       INTEGER NOT NULL,
    id_mekanik         INTEGER NOT NULL,
    jumlah_sparepart   INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_by VARCHAR(100), 
    modified_at TIMESTAMP, 
    modified_by VARCHAR(100),
    CONSTRAINT fk_detail_servis FOREIGN KEY (id_servis)
        REFERENCES servis (id_servis)
        ON DELETE CASCADE,
    CONSTRAINT fk_detail_sparepart FOREIGN KEY (id_sparepart)
        REFERENCES sparepart (id_sparepart)
        ON DELETE CASCADE,
    CONSTRAINT fk_detail_mekanik FOREIGN KEY (id_mekanik)
        REFERENCES mekanik (id_mekanik)
        ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS detail;

-- +migrate StatementEnd
