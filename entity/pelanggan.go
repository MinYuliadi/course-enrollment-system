package entity

import "time"

type Pelanggan struct {
	IDPelanggan   int        `json:"id_pelanggan"`
	NamaPelanggan string     `json:"nama_pelanggan"`
	Alamat        string     `json:"alamat"`
	TelpPelanggan string     `json:"telp_pelanggan"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     string     `json:"created_by"`
	ModifiedAt    *time.Time `json:"modified_at"` // pointer allows NULL
	ModifiedBy    *string    `json:"modified_by"` // pointer allows NULL
}
