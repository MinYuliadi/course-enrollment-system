package entity

import "time"

type Invoice struct {
	IDInvoice        int        `json:"id_invoice"`
	TotalBiaya       int        `json:"total_biaya"`
	TglPembayaran    time.Time  `json:"tgl_pembayaran"`
	MetodePembayaran string     `json:"metode_pembayaran"`
	StatusPembayaran string     `json:"status_pembayaran"`
	IDServis         int        `json:"id_servis"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	ModifiedAt       *time.Time `json:"modified_at"` // pointer allows NULL
	ModifiedBy       *string    `json:"modified_by"` // pointer allows NULL
}
