package entity

import "time"

type Detail struct {
	IDPengerjaan    int        `json:"id_pengerjaan"`
	IDServis        int        `json:"id_servis"`
	IDSparepart     int        `json:"id_sparepart"`
	IDMekanik       int        `json:"id_mekanik"`
	JumlahSparepart int        `json:"jumlah_sparepart"`
	CreatedAt       time.Time  `json:"created_at"`
	CreatedBy       string     `json:"created_by"`
	ModifiedAt      *time.Time `json:"modified_at"` // pointer allows NULL
	ModifiedBy      *string    `json:"modified_by"` // pointer allows NULL
}
