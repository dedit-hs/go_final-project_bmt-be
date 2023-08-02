package controllers

import (
	"net/http"
	"strconv"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/helpers"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file models.File) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file models.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := helpers.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func AddDokumenAnggota(e echo.Context) error {
	anggotaIdValue := e.FormValue("id_anggota")
	dokumenIdValue := e.FormValue("id_dokumen")

	formHeader, err := e.FormFile("file")
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			models.BaseResponse{
				Message: "error",
				Data:    map[string]interface{}{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			models.BaseResponse{
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
	}

	var dokumenAnggota models.DokumenAnggota
	var newDokumenAnggota models.DokumenAnggota
	anggotaId, _ := strconv.Atoi(anggotaIdValue)
	dokumenId, _ := strconv.Atoi(dokumenIdValue)
	newDokumenAnggota.IdAnggota = anggotaId
	newDokumenAnggota.IdDokumen = dokumenId

	cekDokumenAnggota := configs.DB.Where("id_dokumen = ? AND id_anggota = ?", &newDokumenAnggota.IdDokumen, &newDokumenAnggota.IdAnggota).First(&dokumenAnggota)

	if cekDokumenAnggota.Error == gorm.ErrRecordNotFound {

		uploadUrl, err := NewMediaUpload().FileUpload(models.File{File: formFile})
		if err != nil {
			return e.JSON(
				http.StatusInternalServerError,
				models.BaseResponse{
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
		}
		newDokumenAnggota.File = uploadUrl
		result := configs.DB.Create(&newDokumenAnggota)
		if result.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
				"error":   result.Error,
			})
		}
		return e.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newDokumenAnggota,
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "dokumen already exists",
	})

}

func DeleteDokumenAnggota(e echo.Context) error {
	anggotaId := e.Param("anggotaId")
	dokumenId := e.Param("dokumenId")

	var dokumenAnggota models.DokumenAnggota
	dokumenAnggotaToDelete := configs.DB.Where("id_dokumen = ? AND id_anggota = ?", dokumenId, anggotaId).First(&dokumenAnggota)

	if dokumenAnggotaToDelete.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Dokumen tidak ditemukan",
		})
	}

	deleteDokumenAnggota := configs.DB.Delete(&dokumenAnggota)
	if deleteDokumenAnggota.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nil,
	})
}
