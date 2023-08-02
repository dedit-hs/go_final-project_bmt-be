package controllers

import (
	"net/http"
	"strconv"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddPembiayaan(e echo.Context) error {
	var newPembiayaan, pembiayaan models.Pembiayaan
	e.Bind(&newPembiayaan)

	cekPembiayaan := configs.DB.Where("id_anggota = ? AND status = ?", &newPembiayaan.IdAnggota, models.APPROVED).First(&pembiayaan)

	if cekPembiayaan.Error == gorm.ErrRecordNotFound {
		newPembiayaan.Status = models.PENDING
		result := configs.DB.Create(&newPembiayaan)
		if result.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
				"error":   result.Error,
			})
		}
		return e.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newPembiayaan,
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "Masih terdapat pembiayaan yang berjalan",
	})
}

func AddPembayaranPembiayaan(e echo.Context) error {
	anggotaIdValue := e.FormValue("id_anggota")
	pembiayaanIdValue := e.FormValue("id_pembiayaan")
	jumlahValue := e.FormValue("jumlah")

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

	var detailPembiayaan models.DetailPembiayaan
	var newPembayaranPembiayaan models.PembayaranPembiayaan

	anggotaId, _ := strconv.Atoi(anggotaIdValue)
	pembiayaanId, _ := strconv.Atoi(pembiayaanIdValue)
	jumlah, _ := strconv.Atoi(jumlahValue)
	newPembayaranPembiayaan.IdAnggota = anggotaId
	newPembayaranPembiayaan.IdPembiayaan = pembiayaanId
	newPembayaranPembiayaan.Status = models.PENDING

	cekDetailPembiayaan := configs.DB.Where("id_pembiayaan = ?", &newPembayaranPembiayaan.IdPembiayaan).First(&detailPembiayaan)

	if cekDetailPembiayaan.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "detail pembiayaan tidak ditemukan",
		})
	}

	newPembayaranPembiayaan.Jumlah = jumlah
	if newPembayaranPembiayaan.Jumlah != detailPembiayaan.CicilanPerbulan {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "jumlah pembayaran tidak sesuai",
		})
	}
	uploadUrl, err := NewMediaUpload().FileUpload(models.File{File: formFile})
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			models.BaseResponse{
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
	}
	newPembayaranPembiayaan.BuktiTransfer = uploadUrl
	savePembayaranPembiayaan := configs.DB.Create(&newPembayaranPembiayaan)
	if savePembayaranPembiayaan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   savePembayaranPembiayaan.Error,
		})
	}

	return e.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success",
		Data:    savePembayaranPembiayaan,
	})
}
