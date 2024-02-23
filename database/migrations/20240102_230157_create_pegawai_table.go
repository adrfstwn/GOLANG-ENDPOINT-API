package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreatePegawaiTable_20240102_230157 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreatePegawaiTable_20240102_230157{}
	m.Created = "20240102_230157"

	migration.Register("CreatePegawaiTable_20240102_230157", m)
}

// Run the migrations
func (m *CreatePegawaiTable_20240102_230157) Up() {
	m.SQL(`
		CREATE TABLE pegawai (
			id SERIAL PRIMARY KEY,
			nama VARCHAR(255),
			id_statusPegawai INT,
			id_jenisKelamin INT,
			id_agama INT,
			alamat VARCHAR(255)
		);
	`)

}

// Reverse the migrations
func (m *CreatePegawaiTable_20240102_230157) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
