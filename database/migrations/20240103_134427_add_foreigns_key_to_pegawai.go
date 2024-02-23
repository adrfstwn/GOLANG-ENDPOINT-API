package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddForeignsKeyToPegawai_20240103_134427 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddForeignsKeyToPegawai_20240103_134427{}
	m.Created = "20240103_134427"

	migration.Register("AddForeignsKeyToPegawai_20240103_134427", m)
}

// Run the migrations
func (m *AddForeignsKeyToPegawai_20240103_134427) Up() {
	// Tambahkan foreign key ke jenis_kelamin

}

// Reverse the migrations
func (m *AddForeignsKeyToPegawai_20240103_134427) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
