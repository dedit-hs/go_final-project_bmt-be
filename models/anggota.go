package models

import (
	"time"
)

type StatusAnggota string

const (
	Aktif    StatusAnggota = "aktif"
	NonAktif StatusAnggota = "non aktif"
)

type Anggota struct {
	Id         int               `json:"id" form:"id" gorm:"primaryKey"`
	Nama       string            `json:"nama" form:"nama"`
	Nik        uint              `json:"nik" form:"nik" gorm:"unique"`
	Alamat     string            `json:"alamat" form:"alamat"`
	NoHp       string            `json:"no_hp" form:"no_hp" gorm:"unique"`
	Email      string            `json:"email" form:"email" gorm:"unique"`
	Password   string            `json:"password" form:"password"`
	Status     StatusAnggota     `json:"status" form:"status"`
	Dokumen    []DokumenAnggota  `json:"dokumen" gorm:"constraint:OnDelete:CASCADE;foreignKey:IdAnggota"`
	Simpanan   []SimpananAnggota `json:"simpanan" gorm:"constraint:OnDelete:CASCADE;foreignKey:IdAnggota"`
	Pembiayaan []Pembiayaan      `json:"pembiayaan" gorm:"constraint:OnDelete:CASCADE;foreignKey:IdAnggota"`
	CreatedAt  time.Time         `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at" form:"updated_at"`
}

type AnggotaResponse struct {
	Id     int           `json:"id"`
	Nama   string        `json:"nama"`
	Nik    string        `json:"nik"`
	Alamat string        `json:"alamat"`
	NoHp   string        `json:"no_hp"`
	Email  string        `json:"email"`
	Status StatusAnggota `json:"status"`
}

type AnggotaLoginResponse struct {
	Nik   uint   `json:"nik"`
	Token string `json:"token"`
}
