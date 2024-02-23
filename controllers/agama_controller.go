package controllers

import (
	"encoding/json"
	"strconv"
	"uas-api-pegawai/models"

	"github.com/astaxie/beego"
)

type AgamaController struct {
	beego.Controller
}

// GetAllAgama retrieves all Agama
func (c *AgamaController) GetAllAgama() {
	agamas, err := models.GetAllAgama()
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get Agama data"}
	} else {
		c.Data["json"] = agamas
	}
	c.ServeJSON()
}

// GetAgamaByID retrieves an Agama by ID
func (c *AgamaController) GetAgamaByID() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	agama, err := models.GetAgamaByID(id)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get Agama by ID"}
	} else if agama == nil {
		c.Data["json"] = map[string]string{"error": "Agama not found"}
	} else {
		c.Data["json"] = agama
	}
	c.ServeJSON()
}

// CreateAgama creates a new Agama record
func (c *AgamaController) CreateAgama() {
	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid input format", "details": err.Error()}
		c.ServeJSON()
		return
	}

	agama, err := models.CreateAgama(jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to create Agama", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "Agama created successfully", "agama": agama}
	}

	c.ServeJSON()
}

// UpdateAgama updates an existing Agama record
func (c *AgamaController) UpdateAgama() {
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

	err = models.UpdateAgama(id, jsonData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to update Agama", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "Agama updated successfully"}
	}

	c.ServeJSON()
}

// DeleteAgama deletes an Agama record by ID
func (c *AgamaController) DeleteAgama() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid ID parameter"}
		c.ServeJSON()
		return
	}

	err = models.DeleteAgama(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Failed to delete Agama", "details": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "Agama deleted successfully"}
	}

	c.ServeJSON()
}
