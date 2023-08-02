package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/middlewares"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func LoginAnggota(e echo.Context) error {
	var anggota models.Anggota
	e.Bind(&anggota)

	if err := configs.DB.Where("nik = ? AND password = ?", anggota.Nik, anggota.Password).First(&anggota).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Login Failed",
			"error":   err.Error(),
		})
	}

	token, err := middlewares.CreateTokenAnggota(anggota.Id, anggota.Nama)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Login Failed",
			"error":   err.Error(),
		})
	}

	anggotaResponse := models.AnggotaLoginResponse{Nik: anggota.Nik, Token: token}

	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    anggotaResponse,
	})
}

func AddAnggota(e echo.Context) error {
	var newAnggota models.Anggota
	var anggota models.Anggota
	e.Bind(&newAnggota)

	cekUsername := configs.DB.Where("nik = ?", &newAnggota.Nik).First(&anggota)

	if cekUsername.Error == gorm.ErrRecordNotFound {
		newAnggota.Status = models.NonAktif
		result := configs.DB.Create(&newAnggota)
		if result.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
				"error":   result.Error,
			})
		}
		return e.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newAnggota,
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "nik already registered",
	})
}

func GetAllAnggota(e echo.Context) error {
	var anggota []models.Anggota
	result := configs.DB.Preload("Dokumen").Preload("Simpanan").Preload("Pembiayaan").Find(&anggota)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    anggota,
	})
}

func GetAnggotaById(e echo.Context) error {
	anggotaId := e.Param("id")
	var anggota []models.Anggota
	result := configs.DB.Preload(clause.Associations).First(&anggota, anggotaId)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    anggota,
	})
}

func UpdateAnggota(e echo.Context) error {
	anggotaId := e.Param("id")
	var updateAnggota models.Anggota
	cekAnggota := configs.DB.First(&updateAnggota, anggotaId)

	if cekAnggota.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Anggota tidak ditemukan.",
			"error":   cekAnggota.Error,
		})
	}

	e.Bind(&updateAnggota)
	result := configs.DB.Save(&updateAnggota)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    updateAnggota,
	})
}

func DeleteAnggota(e echo.Context) error {
	anggotaId := e.Param("id")

	var anggota models.Anggota
	anggotaToDelete := configs.DB.First(&anggota, anggotaId)

	if anggotaToDelete.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	deleteAnggota := configs.DB.Delete(&anggota)
	if deleteAnggota.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nil,
	})
}
