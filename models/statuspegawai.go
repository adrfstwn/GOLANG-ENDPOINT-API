package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type StatusPegawai struct {
	Id   int    `orm:"auto"`
	Nama string `orm:"column(status)"`
}

func init() {
	//orm.RegisterModel(new(StatusPegawai))
}

func GetAllStatusPegawai() ([]*StatusPegawai, error) {
	o := orm.NewOrm()

	var statusPegawais []*StatusPegawai
	_, err := o.QueryTable("status_pegawai").All(&statusPegawais)
	if err != nil {
		return nil, err
	}

	return statusPegawais, nil
}

func GetStatusPegawaiByID(id int) (*StatusPegawai, error) {
	o := orm.NewOrm()
	statusPegawai := &StatusPegawai{Id: id}
	err := o.Read(statusPegawai)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return statusPegawai, nil
}

func CreateStatusPegawai(data map[string]interface{}) (*StatusPegawai, error) {
	o := orm.NewOrm()

	var statusPegawai StatusPegawai

	// Set values from the provided map
	if nama, ok := data["nama"].(string); ok {
		statusPegawai.Nama = nama
	} else {
		return nil, errors.New("invalid or missing 'nama' field")
	}

	// Insert
	id, err := o.Insert(&statusPegawai)
	if err == nil {
		// Successfully inserted
		statusPegawai.Id = int(id)
		return &statusPegawai, nil
	}

	return nil, err
}

func UpdateStatusPegawai(id int, data map[string]interface{}) error {
	o := orm.NewOrm()

	statusPegawai := StatusPegawai{Id: id}

	if err := o.Read(&statusPegawai); err != nil {
		return err
	}

	// Update fields from the provided map
	if nama, ok := data["nama"].(string); ok {
		statusPegawai.Nama = nama
	} else {
		return errors.New("invalid or missing 'nama' field")
	}

	// Update
	_, err := o.Update(&statusPegawai)
	return err
}

func DeleteStatusPegawai(id int) error {
	o := orm.NewOrm()

	statusPegawai := StatusPegawai{Id: id}

	_, err := o.Delete(&statusPegawai)
	return err
}
