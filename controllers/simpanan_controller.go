package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddSimpanan(e echo.Context) error {
	var newSimpanan models.Simpanan
	var simpanan models.Simpanan
	e.Bind(&newSimpanan)

	cekSimpanan := configs.DB.Where("id = ?", &newSimpanan.Id).First(&simpanan)

	if cekSimpanan.Error == gorm.ErrRecordNotFound {
		result := configs.DB.Create(&newSimpanan)
		if result.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
				"error":   result.Error,
			})
		}
		return e.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newSimpanan,
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "simpanan already exists",
	})
}

func GetAllSimpanan(e echo.Context) error {
	var simpanan []models.Simpanan
	result := configs.DB.Find(&simpanan)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    simpanan,
	})
}

func UpdateSimpanan(e echo.Context) error {
	simpananId := e.Param("id")
	var updateSimpanan models.Simpanan
	cekSimpanan := configs.DB.First(&updateSimpanan, simpananId)

	if cekSimpanan.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Simpanan tidak ditemukan.",
			"error":   cekSimpanan.Error,
		})
	}

	e.Bind(&updateSimpanan)
	result := configs.DB.Save(&updateSimpanan)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    updateSimpanan,
	})
}

func DeleteSimpanan(e echo.Context) error {
	simpananId := e.Param("id")

	var simpanan models.Simpanan
	simpananToDelete := configs.DB.First(&simpanan, simpananId)

	if simpananToDelete.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	deleteSimpanan := configs.DB.Delete(&simpanan)
	if deleteSimpanan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nil,
	})
}
