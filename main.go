package main

import (
	"database/sql"
	"fmt"

	_ "uas-api-pegawai/controllers"
	_ "uas-api-pegawai/models"
	_ "uas-api-pegawai/routers"

	_ "github.com/lib/pq" // Import driver PostgreSQL

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.RunMode = beego.AppConfig.String("runmode")
	beego.BConfig.Listen.HTTPPort, _ = beego.AppConfig.Int("httpport")

	// Konfigurasi koneksi database
	dbConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbname"),
		beego.AppConfig.String("dbhost"),
		beego.AppConfig.String("dbport"),
		beego.AppConfig.String("dbsslmode"),
	)
	// Membuat koneksi database
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		beego.Error("Failed to open database connection:", err)
		return
	}
	defer db.Close()
	// Menguji koneksi ke database
	err = db.Ping()
	if err != nil {
		beego.Error("Failed to ping database:", err)
		return
	}
	// Koneksi database berhasil
	fmt.Println("Connected to the database!")

	//===========================RUN BEEGO=============================//
	beego.Run()
}
