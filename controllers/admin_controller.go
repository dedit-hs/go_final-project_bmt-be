package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_final-project_bmt-be/configs"
	"github.com/dedit-hs/go_final-project_bmt-be/middlewares"
	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginAdminController(e echo.Context) error {
	var admin models.Admin
	e.Bind(&admin)

	if err := configs.DB.Where("username = ? AND password = ?", admin.Username, admin.Password).First(&admin).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Login Failed",
			"error":   err.Error(),
		})
	}

	token, err := middlewares.CreateTokenAdmin(admin.Id, admin.Username)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Login Failed",
			"error":   err.Error(),
		})
	}

	adminResponse := models.AdminLoginResponse{Id: admin.Id, Username: admin.Username, Token: token}

	return e.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    adminResponse,
	})
}

func AddAdminController(e echo.Context) error {
	var newAdmin models.Admin
	var admin models.Admin
	e.Bind(&newAdmin)

	cekUsername := configs.DB.Where("username = ?", &newAdmin.Username).First(&admin)

	if cekUsername.Error == gorm.ErrRecordNotFound {
		result := configs.DB.Create(&newAdmin)
		if result.Error != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal server error",
				"error":   result.Error,
			})
		}
		return e.JSON(http.StatusCreated, models.BaseResponse{
			Message: "Success",
			Data:    newAdmin,
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "username already registered",
	})
}

func ApprovalAnggota(e echo.Context) error {
	idAnggota := e.Param("id")
	var anggota models.Anggota

	cekAnggota := configs.DB.First(&anggota, idAnggota)

	if cekAnggota.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "anggota tidak ditemukan.",
		})

	}

	result := cekAnggota.Update("status", models.Aktif)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   result.Error,
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func ApprovalPembayaranSimpanan(e echo.Context) error {
	idPembayaran := e.Param("id")

	var simpananAnggota models.SimpananAnggota
	var pembayaranSimpanan models.PembayaranSimpanan

	cekPembayaran := configs.DB.First(&pembayaranSimpanan, idPembayaran)
	if cekPembayaran.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "pembayaran simpanan tidak ditemukan",
		})
	}
	cekSimpananAnggota := configs.DB.Where("id_anggota = ? AND id_simpanan = ?", pembayaranSimpanan.IdAnggota, pembayaranSimpanan.IdSimpanan).First(&simpananAnggota)
	if cekSimpananAnggota.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "simpanan tidak ditemukan",
		})
	}
	newSaldo := simpananAnggota.Saldo + pembayaranSimpanan.Jumlah
	updateSaldoSimpanan := configs.DB.Where("id_anggota = ? AND id_simpanan = ?", pembayaranSimpanan.IdAnggota, pembayaranSimpanan.IdSimpanan).First(&simpananAnggota).Update("saldo", newSaldo)

	if updateSaldoSimpanan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   updateSaldoSimpanan.Error,
		})
	}

	updateStatusPembayaran := configs.DB.First(&pembayaranSimpanan, idPembayaran).Update("status", models.APPROVED)

	if updateStatusPembayaran.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   updateStatusPembayaran.Error,
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func RejectPembayaranSimpanan(e echo.Context) error {
	idPembayaran := e.Param("id")

	var pembayaranSimpanan models.PembayaranSimpanan

	cekPembayaran := configs.DB.First(&pembayaranSimpanan, idPembayaran)
	if cekPembayaran.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "pembayaran simpanan tidak ditemukan",
		})
	}

	updateStatusPembayaran := configs.DB.First(&pembayaranSimpanan, idPembayaran).Update("status", models.REJECTED)

	if updateStatusPembayaran.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   updateStatusPembayaran.Error,
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func ApprovalPembiayaan(e echo.Context) error {
	idPembiayaan := e.Param("id")
	var pembiayaan models.Pembiayaan
	var detailPembiayaan models.DetailPembiayaan

	cekPembiayaan := configs.DB.First(&pembiayaan, idPembiayaan)
	if cekPembiayaan.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "pembiayaan tidak ditemukan",
		})
	}

	if pembiayaan.Status == models.APPROVED {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "pembiayaan already approved",
		})
	}

	updateStatusPembiayaan := configs.DB.First(&pembiayaan, idPembiayaan).UpdateColumn("status", models.APPROVED)
	if updateStatusPembiayaan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	detailPembiayaan.IdPembiayaan = pembiayaan.Id
	detailPembiayaan.HargaJual = pembiayaan.Harga + int(float64(pembiayaan.Harga)*(float64(30)/float64(100)))
	detailPembiayaan.Dp = int(float64(detailPembiayaan.HargaJual) * (float64(30) / float64(100)))
	detailPembiayaan.Tenor = 12
	detailPembiayaan.CicilanPerbulan = (detailPembiayaan.HargaJual - detailPembiayaan.Dp) / 12
	detailPembiayaan.SisaCicilan = detailPembiayaan.CicilanPerbulan * detailPembiayaan.Tenor

	saveDetailPembiayaan := configs.DB.Create(&detailPembiayaan)
	if saveDetailPembiayaan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func RejectPembiayaan(e echo.Context) error {
	idPembiayaan := e.Param("id")
	var pembiayaan models.Pembiayaan

	cekPembiayaan := configs.DB.First(&pembiayaan, idPembiayaan)
	if cekPembiayaan.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "pembiayaan tidak ditemukan",
		})
	}

	updateStatusPembiayaan := configs.DB.First(&pembiayaan, idPembiayaan).UpdateColumn("status", models.REJECTED)
	if updateStatusPembiayaan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func ApprovalPembayaranPembiayaan(e echo.Context) error {
	idPembayaran := e.Param("id")

	var pembayaranPembiayaan models.PembayaranPembiayaan

	cekPembayaran := configs.DB.First(&pembayaranPembiayaan, idPembayaran)
	if cekPembayaran.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "pembayaran pembiayaan tidak ditemukan",
		})
	}

	var detailPembiayaan models.DetailPembiayaan
	cekDetailPembiayaan := configs.DB.Where("id_pembiayaan = ?", pembayaranPembiayaan.IdPembiayaan).First(&detailPembiayaan)
	if cekDetailPembiayaan.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "detail pembiayaan tidak ditemukan",
		})
	}

	updateStatusPembayaran := configs.DB.First(&pembayaranPembiayaan, idPembayaran).Update("status", models.APPROVED)

	if updateStatusPembayaran.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   updateStatusPembayaran.Error,
		})
	}

	newSisaCicilan := detailPembiayaan.SisaCicilan - pembayaranPembiayaan.Jumlah
	updateSisaCicilan := configs.DB.Where("id_pembiayaan = ?", pembayaranPembiayaan.IdPembiayaan).First(&detailPembiayaan).Update("sisa_cicilan", newSisaCicilan)

	if updateSisaCicilan.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   updateSisaCicilan.Error,
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func RejectPembayaranPembiayaan(e echo.Context) error {
	idPembayaran := e.Param("id")

	var pembayaranPembiayaan models.PembayaranPembiayaan

	cekPembayaran := configs.DB.First(&pembayaranPembiayaan, idPembayaran)
	if cekPembayaran.Error == gorm.ErrRecordNotFound {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "pembayaran tidak ditemukan",
		})
	}

	updateStatusPembayaran := configs.DB.First(&pembayaranPembiayaan, idPembayaran).Update("status", models.REJECTED)

	if updateStatusPembayaran.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"error":   updateStatusPembayaran.Error,
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
