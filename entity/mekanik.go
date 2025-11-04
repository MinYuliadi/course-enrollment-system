package entity

import "time"

type Mekanik struct {
	IDMekanik   int        `json:"id_mekanik"`
	NamaMekanik string     `json:"nama_mekanik"`
	TelpMekanik string     `json:"telp_mekanik"`
	Keahlian    string     `json:"keahlian"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	ModifiedAt  *time.Time `json:"modified_at"` // pointer allows NULL
	ModifiedBy  *string    `json:"modified_by"` // pointer allows NULL
}
