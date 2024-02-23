package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"uas-api-pegawai/models"

	"github.com/astaxie/beego"
)

type StatusPegawaiController struct {
	beego.Controller
}

// GetAllStatusPegawai retrieves all StatusPegawai
func (c *StatusPegawaiController) GetAllStatusPegawai() {
	statusPegawais, err := models.GetAllStatusPegawai()
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get StatusPegawai data"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = statusPegawais
	c.ServeJSON()
}

// GetStatusPegawaiByID retrieves StatusPegawai by ID
func (c *StatusPegawaiController) GetStatusPegawaiByID() {
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println("idStr:", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	statusPegawai, err := models.GetStatusPegawaiByID(int(id))
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	if statusPegawai == nil {
		c.Data["json"] = map[string]string{"error": "StatusPegawai not found"}
	} else {
		c.Data["json"] = statusPegawai
	}

	c.ServeJSON()
}

// CreateStatusPegawai creates a new StatusPegawai
func (c *StatusPegawaiController) CreateStatusPegawai() {
	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid input format", "details": err.Error()}
		c.ServeJSON()
		return
	}

	statusPegawai, err := models.CreateStatusPegawai(jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to create StatusPegawai", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "StatusPegawai created successfully", "status_pegawai": statusPegawai}
	}

	c.ServeJSON()
}

// UpdateStatusPegawai updates an existing StatusPegawai
func (c *StatusPegawaiController) UpdateStatusPegawai() {
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

	err = models.UpdateStatusPegawai(id, jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to update StatusPegawai", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "StatusPegawai updated successfully"}
	}

	c.ServeJSON()
}

// DeleteStatusPegawai deletes StatusPegawai by ID
func (c *StatusPegawaiController) DeleteStatusPegawai() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	err = models.DeleteStatusPegawai(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to delete StatusPegawai", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "StatusPegawai deleted successfully"}
	}

	c.ServeJSON()
}
