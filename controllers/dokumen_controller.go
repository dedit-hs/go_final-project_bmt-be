package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddDokumen(e echo.Context) error {
	var newDokumen models.Dokumen
	var dokumen models.Dokumen
	e.Bind(&newDokumen)

	cekDokumen := configs.DB.Where("id = ?", &newDokumen.Id).First(&dokumen)

	if cekDokumen.Error == gorm.ErrRecordNotFound {
		result := configs.DB.Create(&newDokumen)
		if result.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
				"error":   result.Error,
			})
		}
		return e.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newDokumen,
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "dokumen already exists",
	})
}

func GetAllDokumen(e echo.Context) error {
	var dokumen []models.Dokumen
	result := configs.DB.Find(&dokumen)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    dokumen,
	})
}

func UpdateDokumen(e echo.Context) error {
	dokumenId := e.Param("id")
	var updateDokumen models.Dokumen
	cekDokumen := configs.DB.First(&updateDokumen, dokumenId)

	if cekDokumen.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Dokumen tidak ditemukan.",
			"error":   cekDokumen.Error,
		})
	}

	e.Bind(&updateDokumen)
	result := configs.DB.Save(&updateDokumen)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    updateDokumen,
	})
}

func DeleteDokumen(e echo.Context) error {
	dokumenId := e.Param("id")

	var dokumen models.Dokumen
	dokumenToDelete := configs.DB.First(&dokumen, dokumenId)

	if dokumenToDelete.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	deleteDokumen := configs.DB.Delete(&dokumen)
	if deleteDokumen.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nil,
	})
}
