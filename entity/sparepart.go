package entity

import "time"

type Sparepart struct {
	IDSparepart   int        `json:"id_sparepart"`
	NamaSparepart string     `json:"nama_sparepart"`
	HargaSatuan   int        `json:"harga_satuan"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     string     `json:"created_by"`
	ModifiedAt    *time.Time `json:"modified_at"` // pointer allows NULL
	ModifiedBy    *string    `json:"modified_by"` // pointer allows NULL
}
