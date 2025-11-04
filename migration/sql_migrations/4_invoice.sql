-- +migrate StatementBegin

-- +migrate Up
CREATE TABLE invoice (
    id_invoice          SERIAL PRIMARY KEY,
    total_biaya         INTEGER,
    tgl_pembayaran      DATE,
    metode_pembayaran   VARCHAR(100),
    status_pembayaran   VARCHAR(100),
    id_servis           INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    created_by VARCHAR(100), 
    modified_at TIMESTAMP, 
    modified_by VARCHAR(100),
    CONSTRAINT fk_invoice_servis FOREIGN KEY (id_servis)
        REFERENCES servis (id_servis)
        ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS invoice;

-- +migrate StatementEnd
