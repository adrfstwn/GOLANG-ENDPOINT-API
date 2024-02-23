package routers

import (
	"uas-api-pegawai/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// Pegawai routes
	beego.Router("/pegawai", &controllers.PegawaiController{}, "get:GetAllPegawai")
	beego.Router("/pegawai/:id", &controllers.PegawaiController{}, "get:GetPegawaiByID")
	beego.Router("/pegawai", &controllers.PegawaiController{}, "post:CreatePegawai")
	beego.Router("/pegawai/:id", &controllers.PegawaiController{}, "put:UpdatePegawai")
	beego.Router("/pegawai/:id", &controllers.PegawaiController{}, "delete:DeletePegawai")

	// Agama routes
	beego.Router("/agama", &controllers.AgamaController{}, "get:GetAllAgama")
	beego.Router("/agama/:id", &controllers.AgamaController{}, "get:GetAgamaByID")
	beego.Router("/agama/", &controllers.AgamaController{}, "post:CreateAgama")
	beego.Router("/agama/:id", &controllers.AgamaController{}, "put:UpdateAgama")
	beego.Router("/agama/:id", &controllers.AgamaController{}, "delete:DeleteAgama")

	// Jenis Kelamin routes
	beego.Router("/jeniskelamin", &controllers.JenisKelaminController{}, "get:GetAllJenisKelamin")
	beego.Router("/jeniskelamin/:id", &controllers.JenisKelaminController{}, "get:GetJenisKelaminByID")
	beego.Router("/jeniskelamin/", &controllers.JenisKelaminController{}, "post:CreateJenisKelamin")
	beego.Router("/jeniskelamin/:id", &controllers.JenisKelaminController{}, "put:UpdateJenisKelamin")
	beego.Router("/jeniskelamin/:id", &controllers.JenisKelaminController{}, "delete:DeleteJenisKelamin")

	// Status Pegawai routes
	beego.Router("/status", &controllers.StatusPegawaiController{}, "get:GetAllStatusPegawai")
	beego.Router("/status/:id", &controllers.StatusPegawaiController{}, "get:GetStatusPegawaiByID")
	beego.Router("/status/", &controllers.StatusPegawaiController{}, "post:CreateStatusPegawai")
	beego.Router("/status/:id", &controllers.StatusPegawaiController{}, "put:UpdateStatusPegawai")
	beego.Router("/status/:id", &controllers.StatusPegawaiController{}, "delete:DeleteStatusPegawai")

}
