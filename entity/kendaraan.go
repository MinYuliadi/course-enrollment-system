package entity

import "time"

type Kendaraan struct {
	IDMotor          int        `json:"id_motor"`
	IDPelanggan      int        `json:"id_pelanggan"`
	PlatMotor        string     `json:"plat_motor"`
	TahunProduksi    int        `json:"tahun_produksi"`
	StatusKepmilikan string     `json:"status_kepmilikan"`
	Merek            string     `json:"merek"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	ModifiedAt       *time.Time `json:"modified_at"` // pointer allows NULL
	ModifiedBy       *string    `json:"modified_by"` // pointer allows NULL
}
