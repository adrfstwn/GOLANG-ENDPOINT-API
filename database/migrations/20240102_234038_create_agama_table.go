package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateAgamaTable_20240102_234038 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAgamaTable_20240102_234038{}
	m.Created = "20240102_234038"

	migration.Register("CreateAgamaTable_20240102_234038", m)
}

// Run the migrations
func (m *CreateAgamaTable_20240102_234038) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`
		CREATE TABLE agama (
			id SERIAL PRIMARY KEY,
			agama VARCHAR(255)
		);
	`)
}

// Reverse the migrations
func (m *CreateAgamaTable_20240102_234038) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
