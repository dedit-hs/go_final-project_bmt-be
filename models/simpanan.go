package models

import "time"

type Simpanan struct {
	Id              int       `json:"id" form:"id"  gorm:"primaryKey"`
	Nama            string    `json:"nama" form:"nama" gorm:"unique"`
	MinimalSimpanan int       `json:"minimal_simpanan" form:"minimal_simpanan"`
	CreatedAt       time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" form:"updated_at"`
}

type SimpananAnggota struct {
	Id         int       `json:"id" form:"id"  gorm:"primaryKey"`
	IdSimpanan int       `json:"id_simpanan" form:"id_simpanan"`
	IdAnggota  int       `json:"id_anggota" form:"id_anggota"`
	Saldo      int       `json:"saldo" form:"saldo"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
}
