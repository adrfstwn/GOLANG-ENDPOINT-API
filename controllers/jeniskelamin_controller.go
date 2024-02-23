package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"uas-api-pegawai/models"

	"github.com/astaxie/beego"
)

type JenisKelaminController struct {
	beego.Controller
}

// GetAllJenisKelamin retrieves all JenisKelamin records
func (c *JenisKelaminController) GetAllJenisKelamin() {
	jenisKelaminData, err := models.GetAllJenisKelamin()
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get JenisKelamin data"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = jenisKelaminData
	c.ServeJSON()
}

// GetJenisKelaminByID retrieves a JenisKelamin by ID
func (c *JenisKelaminController) GetJenisKelaminByID() {
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println("idStr:", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	jenisKelamin, err := models.GetJenisKelaminByID(id)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	if jenisKelamin == nil {
		c.Data["json"] = map[string]string{"error": "JenisKelamin not found"}
	} else {
		c.Data["json"] = jenisKelamin
	}

	c.ServeJSON()
}

// CreateJenisKelamin creates a new JenisKelamin record
func (c *JenisKelaminController) CreateJenisKelamin() {
	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid input format", "details": err.Error()}
		c.ServeJSON()
		return
	}

	jenisKelamin, err := models.CreateJenisKelamin(jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to create JenisKelamin", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "JenisKelamin created successfully", "data": jenisKelamin}
	}

	c.ServeJSON()
}

// UpdateJenisKelamin updates an existing JenisKelamin record
func (c *JenisKelaminController) UpdateJenisKelamin() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	err = decoder.Decode(&jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid input format", "details": err.Error()}
		c.ServeJSON()
		return
	}

	err = models.UpdateJenisKelamin(id, jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to update JenisKelamin", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "JenisKelamin updated successfully"}
	}

	c.ServeJSON()
}

// DeleteJenisKelamin deletes a JenisKelamin record by ID
func (c *JenisKelaminController) DeleteJenisKelamin() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	err = models.DeleteJenisKelamin(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to delete JenisKelamin", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "JenisKelamin deleted successfully"}
	}

	c.ServeJSON()
}
