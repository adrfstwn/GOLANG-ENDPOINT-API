package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type JenisKelamin struct {
	Id   int    `orm:"auto"`
	Nama string `orm:"column(jenis_kelamin)"`
}

func init() {
	//orm.RegisterModel(new(JenisKelamin))
}

func GetAllJenisKelamin() ([]*JenisKelamin, error) {
	o := orm.NewOrm()

	var jenisKelamins []*JenisKelamin
	_, err := o.QueryTable("jenis_kelamin").All(&jenisKelamins)
	if err != nil {
		return nil, err
	}

	return jenisKelamins, nil
}

func GetJenisKelaminByID(id int) (*JenisKelamin, error) {
	o := orm.NewOrm()
	jenisKelamin := &JenisKelamin{Id: id}
	err := o.Read(jenisKelamin)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return jenisKelamin, nil
}

func CreateJenisKelamin(data map[string]interface{}) (*JenisKelamin, error) {
	o := orm.NewOrm()

	var jenisKelamin JenisKelamin

	// Set values from the provided map
	if nama, ok := data["nama"].(string); ok {
		jenisKelamin.Nama = nama
	} else {
		return nil, errors.New("invalid or missing 'nama' field")
	}

	// Insert
	id, err := o.Insert(&jenisKelamin)
	if err == nil {
		// Successfully inserted
		jenisKelamin.Id = int(id)
		return &jenisKelamin, nil
	}

	return nil, err
}

func UpdateJenisKelamin(id int, data map[string]interface{}) error {
	o := orm.NewOrm()

	jenisKelamin := JenisKelamin{Id: id}

	if err := o.Read(&jenisKelamin); err != nil {
		return err
	}

	// Update fields from the provided map
	if nama, ok := data["nama"].(string); ok {
		jenisKelamin.Nama = nama
	} else {
		return errors.New("invalid or missing 'nama' field")
	}

	// Update
	_, err := o.Update(&jenisKelamin)
	return err
}

func DeleteJenisKelamin(id int) error {
	o := orm.NewOrm()

	jenisKelamin := JenisKelamin{Id: id}

	_, err := o.Delete(&jenisKelamin)
	return err
}
