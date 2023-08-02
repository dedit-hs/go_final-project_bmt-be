package models

import "time"

type StatusPembayaran string

var (
	APPROVED StatusPembayaran = "approved"
	PENDING  StatusPembayaran = "pending"
	REJECTED StatusPembayaran = "rejected"
	END      StatusPembayaran = "end"
)

type PembayaranPembiayaan struct {
	Id            int              `json:"id" form:"id"  gorm:"primaryKey"`
	IdAnggota     int              `json:"id_anggota" form:"id_anggota"`
	IdPembiayaan  int              `json:"id_pembiayaan" form:"id_pembiayaan"`
	Jumlah        int              `json:"jumlah" form:"jumlah"`
	BuktiTransfer string           `json:"bukti_transfer" form:"bukti_transfer"`
	Status        StatusPembayaran `json:"status" form:"status"`
	CreatedAt     time.Time        `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at" form:"updated_at"`
}

type PembayaranSimpanan struct {
	Id            int              `json:"id" form:"id"  gorm:"primaryKey"`
	IdAnggota     int              `json:"id_anggota" form:"id_anggota"`
	IdSimpanan    int              `json:"id_simpanan" form:"id_simpanan"`
	Jumlah        int              `json:"jumlah" form:"jumlah"`
	BuktiTransfer string           `json:"bukti_transfer" form:"bukti_transfer"`
	Status        StatusPembayaran `json:"status" form:"status"`
	CreatedAt     time.Time        `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at" form:"updated_at"`
}
