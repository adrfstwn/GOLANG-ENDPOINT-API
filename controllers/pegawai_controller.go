package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"uas-api-pegawai/models"

	"github.com/astaxie/beego"
)

type PegawaiController struct {
	beego.Controller
}
type Pegawai struct {
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	JenisKelamin struct {
		ID   int    `json:"id"`
		Nama string `json:"jenis_kelamin"`
	} `json:"jenis_kelamin"`
	Agama struct {
		ID   int    `json:"id"`
		Nama string `json:"agama"`
	} `json:"agama"`
	Status struct {
		ID   int    `json:"id"`
		Nama string `json:"status"`
	} `json:"status"`
}

// Define a custom response structure
type PegawaiResponse struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	JenisKelamin  JenisKelaminResponse
	StatusPegawai StatusPegawaiResponse
	Agama         AgamaResponse
}

type JenisKelaminResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"jenis_kelamin"`
}

type StatusPegawaiResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"status"`
}

type AgamaResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"agama"`
}

// mengambil semua pegawai
func (c *PegawaiController) GetAllPegawai() {
	query := make(map[string]string)
	fields := make([]string, 0)
	sortby := make([]string, 0)
	order := make([]string, 0)
	offset := int64(0)
	limit := int64(0)

	pegawaiData, err := models.GetAllPegawai(query, fields, sortby, order, offset, limit)
	if err != nil {
		beego.Error("Error getting Pegawai data:", err)
		c.Data["json"] = map[string]string{"error": "Failed to get Pegawai data"}
		c.ServeJSON()
		return
	}

	//mengecek agar nilai tidak nill
	if data, ok := pegawaiData["data"].([]models.Pegawai); ok {
		c.Data["json"] = data
	} else {
		c.Data["json"] = map[string]string{"error": "Invalid Pegawai data format"}
	}

	c.ServeJSON()
}

// GetPegawaiByID retrieves a Pegawai by ID
func (c *PegawaiController) GetPegawaiByID() {
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println("idStr:", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	pegawai, err := models.GetPegawaiById(int(id))
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	if pegawai == nil {
		c.Data["json"] = map[string]string{"error": "Pegawai not found"}
	} else {
		// Create a custom response structure with the desired order
		response := PegawaiResponse{
			ID:     pegawai.Id,
			Nama:   pegawai.Nama,
			Alamat: pegawai.Alamat,
			JenisKelamin: JenisKelaminResponse{
				ID:   pegawai.JenisKelamin.Id,
				Nama: pegawai.JenisKelamin.Nama,
			},
			StatusPegawai: StatusPegawaiResponse{
				ID:   pegawai.StatusPegawai.Id,
				Nama: pegawai.StatusPegawai.Nama,
			},
			Agama: AgamaResponse{
				ID:   pegawai.Agama.Id,
				Nama: pegawai.Agama.Agama,
			},
		}

		// Return the custom response
		c.Data["json"] = response
	}

	c.ServeJSON()
}

func (c *PegawaiController) CreatePegawai() {
	//fmt.Println("Request Headers:", c.Ctx.Request.Header)

	// Parse request body to map
	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		c.Data["json"] = map[string]interface{}{"error": "Invalid input format", "details": err.Error()}
		c.ServeJSON()
		return
	}

	fmt.Println("Received JSON data:", jsonData)

	// Create Pegawai using the model function
	pegawai, err := models.CreatePegawai(jsonData)
	if err != nil {
		fmt.Println("Error creating Pegawai:", err)
		c.Data["json"] = map[string]interface{}{"error": "Failed to create Pegawai", "details": err.Error()}
	} else {
		// Respond with success and include the created Pegawai details
		c.Data["json"] = map[string]interface{}{"success": true, "message": "Pegawai created successfully", "pegawai": pegawai}
	}

	c.ServeJSON()
}

func (c *PegawaiController) UpdatePegawai() {
	// Parse request body to map
	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid input format", "details": err.Error()}
		c.ServeJSON()
		return
	}

	// Get Pegawai ID from URL parameter
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	// Update Pegawai
	updatedPegawai, err := models.UpdatePegawai(id, jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to update Pegawai", "details": err.Error()}
		c.ServeJSON()
		return
	}

	// Respond with success and include the updated Pegawai details
	c.Data["json"] = map[string]interface{}{"success": true, "message": "Pegawai updated successfully", "pegawai": updatedPegawai}
	c.ServeJSON()
}

// Add a new method for deleting Pegawai by ID
func (c *PegawaiController) DeletePegawai() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	err = models.DeletePegawai(int(id))
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"success": "Pegawai deleted successfully"}
	}

	c.ServeJSON()
}
