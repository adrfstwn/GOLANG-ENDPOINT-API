package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreatedStatuspegawaiTable_20240103_134804 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreatedStatuspegawaiTable_20240103_134804{}
	m.Created = "20240103_134804"

	migration.Register("CreatedStatuspegawaiTable_20240103_134804", m)
}

// Run the migrations
func (m *CreatedStatuspegawaiTable_20240103_134804) Up() {
	m.SQL(`
		CREATE TABLE statuspegawai (
			id SERIAL PRIMARY KEY,
			status VARCHAR(255)
			
		);
	`)

}

// Reverse the migrations
func (m *CreatedStatuspegawaiTable_20240103_134804) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
