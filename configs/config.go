package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/dedit-hs/go_final-project_bmt-be/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Fail to load env file.")
	}
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: os.Getenv("FPDB_Username"),
		DB_Password: os.Getenv("FPDB_Password"),
		DB_Port:     os.Getenv("FPDB_Port"),
		DB_Host:     os.Getenv("FPDB_Host"),
		DB_Name:     os.Getenv("FPDB_Name"),
	}
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	migrateDB()
}

func migrateDB() {
	DB.AutoMigrate(&models.Admin{}, &models.Simpanan{}, &models.Dokumen{}, &models.ObjekPembiayaan{}, &models.PembayaranPembiayaan{}, &models.PembayaranSimpanan{}, &models.Anggota{}, &models.Pembiayaan{}, &models.DetailPembiayaan{}, &models.DokumenAnggota{}, &models.SimpananAnggota{})
}

func EnvCloudName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
