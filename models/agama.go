package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type Agama struct {
	Id    int    `json:"id"`
	Agama string `json:"agama"`
}

func init() {
	// orm.RegisterModel(new(Agama))
}

func GetAllAgama() ([]*Agama, error) {
	o := orm.NewOrm()

	var agamas []*Agama
	_, err := o.QueryTable("agama").All(&agamas)
	if err != nil {
		return nil, err
	}

	return agamas, nil
}

func GetAgamaByID(id int) (*Agama, error) {
	o := orm.NewOrm()
	agama := &Agama{Id: id}
	err := o.Read(agama)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return agama, nil
}

// CreateAgama creates a new Agama record
func CreateAgama(data map[string]interface{}) (*Agama, error) {
	o := orm.NewOrm()

	var agama Agama

	// Set values from the provided map
	if agamaName, ok := data["agama"].(string); ok {
		agama.Agama = agamaName
	} else {
		return nil, errors.New("invalid or missing 'agama' field")
	}

	// Insert
	id, err := o.Insert(&agama)
	if err == nil {
		// Successfully inserted
		agama.Id = int(id)
		return &agama, nil
	}

	return nil, err
}

// UpdateAgama updates an existing Agama record
func UpdateAgama(id int, data map[string]interface{}) error {
	o := orm.NewOrm()

	agama := Agama{Id: id}

	if err := o.Read(&agama); err != nil {
		return err
	}

	// Update fields from the provided map
	if agamaName, ok := data["agama"].(string); ok {
		agama.Agama = agamaName
	} else {
		return errors.New("invalid or missing 'agama' field")
	}

	// Update
	_, err := o.Update(&agama)
	return err
}

// DeleteAgama deletes an Agama record by ID
func DeleteAgama(id int) error {
	o := orm.NewOrm()

	agama := Agama{Id: id}

	_, err := o.Delete(&agama)
	return err
}
