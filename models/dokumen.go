package models

import (
	"mime/multipart"
	"time"
)

type Dokumen struct {
	Id        int       `json:"id" form:"id"  gorm:"primaryKey"`
	Nama      string    `json:"nama" form:"nama" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type DokumenAnggota struct {
	Id        int       `json:"id" form:"id"  gorm:"primaryKey"`
	IdDokumen int       `json:"id_dokumen" form:"id_dokumen"`
	IdAnggota int       `json:"id_anggota" form:"id_anggota"`
	File      string    `json:"file" form:"file"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}
