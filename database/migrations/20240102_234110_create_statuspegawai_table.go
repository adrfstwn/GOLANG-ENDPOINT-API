package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateStatuspegawaiTable_20240102_234110 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateStatuspegawaiTable_20240102_234110{}
	m.Created = "20240102_234110"

	migration.Register("CreateStatuspegawaiTable_20240102_234110", m)
}

// Run the migrations
func (m *CreateStatuspegawaiTable_20240102_234110) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreateStatuspegawaiTable_20240102_234110) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`
		CREATE TABLE statuspegawai (
			id SERIAL PRIMARY KEY,
			status VARCHAR(255)
			
		);
	`)
}
