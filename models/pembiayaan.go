package models

import "time"

type ObjekPembiayaan struct {
	Id        int       `json:"id" form:"id"  gorm:"primaryKey"`
	Nama      string    `json:"nama" form:"nama" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type DetailPembiayaan struct {
	Id              int       `json:"id" form:"id"  gorm:"primaryKey"`
	IdPembiayaan    int       `json:"id_pembiayaan" form:"id_pembiayaan"`
	HargaJual       int       `json:"harga_jual" form:"harga_jual"`
	Dp              int       `json:"dp" form:"dp"`
	Tenor           int       `json:"tenor" form:"tenor"`
	CicilanPerbulan int       `json:"cicilan_perbulan" form:"cicilan_perbulan"`
	SisaCicilan     int       `json:"sisa_cicilan" form:"sisa_cicilan"`
	CreatedAt       time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" form:"updated_at"`
}

type Pembiayaan struct {
	Id        int              `json:"id" form:"id"  gorm:"primaryKey"`
	IdAnggota int              `json:"id_anggota" form:"id_anggota"`
	IdObjek   int              `json:"id_objek" form:"id_objek"`
	Merk      string           `json:"merk" form:"merk"`
	Tipe      string           `json:"tipe" form:"tipe"`
	Harga     int              `json:"harga" form:"harga"`
	Detail    DetailPembiayaan `json:"detail" gorm:"constraint:OnDelete:CASCADE;foreignKey:IdPembiayaan"`
	Status    StatusPembayaran `json:"status" form:"status"`
	CreatedAt time.Time        `json:"created_at" form:"created_at"`
	UpdatedAt time.Time        `json:"updated_at" form:"updated_at"`
}
