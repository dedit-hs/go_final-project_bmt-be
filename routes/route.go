package routes

import (
	"os"

	"github.com/dedit-hs/go_final-project_bmt-be/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/admin/login", controllers.LoginAdminController)
	e.POST("/users/login", controllers.LoginAnggota)
	e.POST("/users", controllers.AddAnggota)

	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	eJWT.PUT("/users/:id", controllers.UpdateAnggota)
	eJWT.GET("/objek", controllers.GetAllObjek)
	eJWT.POST("/dokumen/upload", controllers.AddDokumenAnggota)
	eJWT.POST("/bayarsimpanan", controllers.AddPembayaranSimpanan)
	eJWT.POST("/pembiayaan", controllers.AddPembiayaan)
	eJWT.POST("/bayarpembiayaan", controllers.AddPembayaranPembiayaan)

	adminJWT := e.Group("/admin")
	adminJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))
	adminJWT.POST("", controllers.AddAdminController)

	adminJWT.GET("/users", controllers.GetAllAnggota)
	adminJWT.GET("/users/:id", controllers.GetAnggotaById)
	adminJWT.PUT("/users/:id/approve", controllers.ApprovalAnggota)
	adminJWT.PUT("/users/:id", controllers.UpdateAnggota)
	adminJWT.DELETE("/users/:id", controllers.DeleteAnggota)

	adminJWT.POST("/simpanan", controllers.AddSimpanan)
	adminJWT.GET("/simpanan", controllers.GetAllSimpanan)
	adminJWT.PUT("/simpanan/:id", controllers.UpdateSimpanan)
	adminJWT.DELETE("/simpanan/:id", controllers.DeleteSimpanan)

	adminJWT.POST("/dokumen", controllers.AddDokumen)
	adminJWT.GET("/dokumen", controllers.GetAllDokumen)
	adminJWT.PUT("/dokumen/:id", controllers.UpdateDokumen)
	adminJWT.DELETE("/dokumen/:id", controllers.DeleteDokumen)

	adminJWT.POST("/objek", controllers.AddObjek)
	adminJWT.GET("/objek", controllers.GetAllObjek)
	adminJWT.PUT("/objek/:id", controllers.UpdateObjek)
	adminJWT.DELETE("/objek/:id", controllers.DeleteObjek)

	adminJWT.POST("/dokumen/upload", controllers.AddDokumenAnggota)
	adminJWT.DELETE("/users/:anggotaId/dokumen/:dokumenId", controllers.DeleteDokumenAnggota)

	adminJWT.POST("/bayarsimpanan", controllers.AddPembayaranSimpanan)
	adminJWT.PUT("/bayarsimpanan/:id/approve", controllers.ApprovalPembayaranSimpanan)

	adminJWT.PUT("/pembiayaan/:id/approve", controllers.ApprovalPembiayaan)
	adminJWT.POST("/bayarpembiayaan", controllers.AddPembayaranPembiayaan)
	adminJWT.PUT("/bayarpembiayaan/:id/approve", controllers.ApprovalPembayaranPembiayaan)
	adminJWT.PUT("/bayarpembiayaan/:id/reject", controllers.RejectPembayaranPembiayaan)

	return e
}
