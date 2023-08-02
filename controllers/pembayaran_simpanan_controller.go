package controllers

import (
	"net/http"
	"strconv"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddPembayaranSimpanan(e echo.Context) error {
	anggotaIdValue := e.FormValue("id_anggota")
	simpananIdValue := e.FormValue("id_simpanan")
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

	var simpananAnggota models.SimpananAnggota
	var newPembayaranSimpanan models.PembayaranSimpanan

	anggotaId, _ := strconv.Atoi(anggotaIdValue)
	simpananId, _ := strconv.Atoi(simpananIdValue)
	jumlah, _ := strconv.Atoi(jumlahValue)
	newPembayaranSimpanan.IdAnggota = anggotaId
	newPembayaranSimpanan.IdSimpanan = simpananId
	newPembayaranSimpanan.Status = models.PENDING

	cekSimpananAnggota := configs.DB.Where("id_simpanan = ? AND id_anggota = ?", &newPembayaranSimpanan.IdSimpanan, &newPembayaranSimpanan.IdAnggota).First(&simpananAnggota)

	if cekSimpananAnggota.Error == gorm.ErrRecordNotFound {
		simpananAnggota.IdAnggota = newPembayaranSimpanan.IdAnggota
		simpananAnggota.IdSimpanan = newPembayaranSimpanan.IdSimpanan
		simpananAnggota.Saldo = 0
		saveSimpananAnggota := configs.DB.Create(&simpananAnggota)
		if saveSimpananAnggota.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error",
				"error":   saveSimpananAnggota.Error,
			})
		}
	}

	var simpanan models.Simpanan
	simpanan.Id = newPembayaranSimpanan.IdSimpanan
	cekMinimal := configs.DB.First(&simpanan, newPembayaranSimpanan.IdSimpanan)
	if cekMinimal.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   cekMinimal.Error,
		})
	}
	if simpanan.Nama == "pokok" && simpananAnggota.Saldo >= simpanan.MinimalSimpanan {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "simpanan pokok hanya bisa dibayarkan sekali",
		})
	}

	newPembayaranSimpanan.Jumlah = jumlah
	if newPembayaranSimpanan.Jumlah != simpanan.MinimalSimpanan {
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
	newPembayaranSimpanan.BuktiTransfer = uploadUrl
	savePembayaranSimpanan := configs.DB.Create(&newPembayaranSimpanan)
	if savePembayaranSimpanan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   savePembayaranSimpanan.Error,
		})
	}

	return e.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success",
		Data:    newPembayaranSimpanan,
	})
}
