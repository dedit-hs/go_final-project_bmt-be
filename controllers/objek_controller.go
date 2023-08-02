package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddObjek(e echo.Context) error {
	var newObjek models.ObjekPembiayaan
	var objek models.ObjekPembiayaan
	e.Bind(&newObjek)

	cekObjek := configs.DB.Where("id = ?", &newObjek.Id).First(&objek)

	if cekObjek.Error == gorm.ErrRecordNotFound {
		result := configs.DB.Create(&newObjek)
		if result.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
				"error":   result.Error,
			})
		}
		return e.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newObjek,
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "objek already exists",
	})
}

func GetAllObjek(e echo.Context) error {
	var objek []models.ObjekPembiayaan
	result := configs.DB.Find(&objek)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    objek,
	})
}

func UpdateObjek(e echo.Context) error {
	objekId := e.Param("id")
	var updateObjek models.ObjekPembiayaan
	cekObjek := configs.DB.First(&updateObjek, objekId)

	if cekObjek.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Objek tidak ditemukan.",
			"error":   cekObjek.Error,
		})
	}

	e.Bind(&updateObjek)
	result := configs.DB.Save(&updateObjek)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    updateObjek,
	})
}

func DeleteObjek(e echo.Context) error {
	objekId := e.Param("id")

	var objek models.ObjekPembiayaan
	objekToDelete := configs.DB.First(&objek, objekId)

	if objekToDelete.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	ObjekLurah := configs.DB.Delete(&objek)
	if ObjekLurah.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nil,
	})
}
