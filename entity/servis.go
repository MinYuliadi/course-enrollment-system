package entity

import "time"

type Servis struct {
	IDServis         int        `json:"id_servis"`
	TglServis        time.Time  `json:"tgl_servis"`
	Keluhan          string     `json:"keluhan"`
	StatusPengerjaan string     `json:"status_pengerjaan"`
	IDMotor          int        `json:"id_motor"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	ModifiedAt       *time.Time `json:"modified_at"` // pointer allows NULL
	ModifiedBy       *string    `json:"modified_by"` // pointer allows NULL
}
