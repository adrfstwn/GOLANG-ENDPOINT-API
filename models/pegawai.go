package models

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Pegawai struct {
	Id            int            `json:"id"`
	Nama          string         `json:"nama"`
	Alamat        string         `json:"alamat"`
	JenisKelamin  *JenisKelamin  `orm:"rel(fk);column(id_jeniskelamin)" json:"jenis_kelamin"`
	Agama         *Agama         `orm:"rel(fk);column(id_agama)" json:"agama"`
	StatusPegawai *StatusPegawai `orm:"rel(fk);column(id_statuspegawai)" json:"status"`
	RegDate       time.Time      `orm:"auto_now_add;type(datetime)" json:"reg_date"`
}

func init() {
	// Register the database driver
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// Register the default database alias named `default`
	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=golang-api-uas host=127.0.0.1 port=5432 sslmode=disable")

	// Register the models
	orm.RegisterModel(new(StatusPegawai))
	orm.RegisterModel(new(JenisKelamin))
	orm.RegisterModel(new(Agama))
	orm.RegisterModel(new(Pegawai))

	// Run the database migration
	orm.RunSyncdb("default", false, false)
}

// get pegawai
func GetAllPegawai(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (result map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Pegawai))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: unused 'order' fields")
		}
	}

	var l []Pegawai
	qs = qs.OrderBy(sortFields...).RelatedSel("JenisKelamin", "Agama", "StatusPegawai")
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		result := map[string]interface{}{"data": l}
		return result, nil
	}
	return nil, err
}

func GetPegawaiById(id int) (v *Pegawai, err error) {
	o := orm.NewOrm()
	v = &Pegawai{Id: int(id)}
	if err = o.Read(v); err == nil {
		// Load related data
		if err := o.Read(v.JenisKelamin); err != nil {
			return nil, err
		}
		if err := o.Read(v.Agama); err != nil {
			return nil, err
		}
		if err := o.Read(v.StatusPegawai); err != nil {
			return nil, err
		}
		return v, nil
	}
	return nil, err
}

func CreatePegawai(data map[string]interface{}) (*Pegawai, error) {
	o := orm.NewOrm()

	var pegawai Pegawai

	// Set values from the provided map
	if nama, ok := data["nama"].(string); ok {
		pegawai.Nama = nama
	} else {
		return nil, errors.New("invalid or missing 'nama' field")
	}

	if alamat, ok := data["alamat"].(string); ok {
		pegawai.Alamat = alamat
	} else {
		return nil, errors.New("invalid or missing 'alamat' field")
	}

	if jenisKelaminID, ok := data["jenis_kelamin"].(float64); ok {
		pegawai.JenisKelamin = &JenisKelamin{Id: int(jenisKelaminID)}
	} else {
		return nil, errors.New("invalid or missing 'jenis_kelamin' field")
	}

	if agamaID, ok := data["agama"].(float64); ok {
		pegawai.Agama = &Agama{Id: int(agamaID)}
	} else {
		return nil, errors.New("invalid or missing 'agama' field")
	}

	if statusID, ok := data["status"].(float64); ok {
		pegawai.StatusPegawai = &StatusPegawai{Id: int(statusID)}
	} else {
		return nil, errors.New("invalid or missing 'status' field")
	}

	pegawai.RegDate = time.Now()

	// Insert
	id, err := o.Insert(&pegawai)
	if err == nil {
		// successfully inserted
		pegawai.Id = int(id)
		return &pegawai, nil
	}

	return nil, err
}

// UpdatePegawai updates the details of an existing Pegawai
func UpdatePegawai(id int, data map[string]interface{}) (*Pegawai, error) {
	o := orm.NewOrm()

	// Retrieve the existing Pegawai by ID
	pegawai, err := GetPegawaiById(id)
	if err != nil {
		return nil, err
	}

	// Update Pegawai details based on the provided map
	if nama, ok := data["nama"].(string); ok {
		pegawai.Nama = nama
	}

	if alamat, ok := data["alamat"].(string); ok {
		pegawai.Alamat = alamat
	}

	if jenisKelaminID, ok := data["jenis_kelamin"].(float64); ok {
		pegawai.JenisKelamin = &JenisKelamin{Id: int(jenisKelaminID)}
	}

	if agamaID, ok := data["agama"].(float64); ok {
		pegawai.Agama = &Agama{Id: int(agamaID)}
	}

	if statusID, ok := data["status"].(float64); ok {
		pegawai.StatusPegawai = &StatusPegawai{Id: int(statusID)}
	}

	// Save the updated Pegawai
	_, err = o.Update(pegawai)
	if err != nil {
		return nil, err
	}

	return pegawai, nil
}

// DeletePegawai deletes a Pegawai record by ID
func DeletePegawai(id int) error {
	o := orm.NewOrm()
	pegawai := &Pegawai{Id: id}

	if o.Read(pegawai) == nil {
		_, err := o.Delete(pegawai)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Pegawai not found")
	}

	return nil
}
