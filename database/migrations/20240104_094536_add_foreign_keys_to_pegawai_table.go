package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddForeignKeysToPegawaiTable_20240104_094536 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddForeignKeysToPegawaiTable_20240104_094536{}
	m.Created = "20240104_094536"

	migration.Register("AddForeignKeysToPegawaiTable_20240104_094536", m)
}

// Run the migrations
func (m *AddForeignKeysToPegawaiTable_20240104_094536) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE pegawai ADD CONSTRAINT fk_pegawai_statuspegawai FOREIGN KEY (id_statuspegawai) REFERENCES status_pegawai(id) ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE pegawai ADD CONSTRAINT fk_pegawai_jeniskelamin FOREIGN KEY (id_jeniskelamin) REFERENCES jenis_kelamin(id) ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE pegawai ADD CONSTRAINT fk_pegawai_agama FOREIGN KEY (id_agama) REFERENCES agama(id) ON DELETE RESTRICT ON UPDATE CASCADE;")

}

// Reverse the migrations
func (m *AddForeignKeysToPegawaiTable_20240104_094536) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
