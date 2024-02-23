package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddForeignsKeyToPegawai_20240103_135146 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddForeignsKeyToPegawai_20240103_135146{}
	m.Created = "20240103_135146"

	migration.Register("AddForeignsKeyToPegawai_20240103_135146", m)
}

// Run the migrations
func (m *AddForeignsKeyToPegawai_20240103_135146) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE pegawai ADD CONSTRAINT fk_pegawai_statuspegawai FOREIGN KEY (id_statuspegawai) REFERENCES statuspegawai(id) ON DELETE RESTRICT ON UPDATE CASCADE;")
}

// Reverse the migrations
func (m *AddForeignsKeyToPegawai_20240103_135146) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
