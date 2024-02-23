package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateJeniskelaminTable_20240102_234055 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateJeniskelaminTable_20240102_234055{}
	m.Created = "20240102_234055"

	migration.Register("CreateJeniskelaminTable_20240102_234055", m)
}

// Run the migrations
func (m *CreateJeniskelaminTable_20240102_234055) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`
		CREATE TABLE jeniskelamin (
			id SERIAL PRIMARY KEY,
			jenis_kelamin VARCHAR(255)
		);
	`)
}

// Reverse the migrations
func (m *CreateJeniskelaminTable_20240102_234055) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
