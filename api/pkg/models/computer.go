package models

import (
	"gorm.io/gorm"
)

type Computer struct {
	gorm.Model
	Name      string `json:"name"`
	Processor string `json:"processor"`
	RAM       string `json:"ram"`
	Storage   string `json:"storage"`
	OS        string `json:"os"`
	User      string `json:"user"`
	Unit      string `json:"unit"`
	Note      string `json:"note,omitempty"`
}

// constructs a new Computer object
func NewComputer() *Computer {
	return &Computer{}
}

// lists all the computers that exists in the database
func (c Computer) List() []Computer {
	var computers []Computer

	DB.Find(&computers)

	return computers
}

// gets one computer from database using the ID
func (c Computer) Get(id string) Computer {
	var computer Computer

	DB.Find(&computer, id)

	return computer
}

// adds a new Computer to database
func (c Computer) Add(computer *Computer) {
	DB.Create(&computer)
}

// updates an existing computer in the database
func (c Computer) Update(computer Computer) {
	DB.Model(&computer).Updates(computer)
}

// deletes a computer from the database
func (c Computer) Delete(id string) Computer {
	var computer Computer

	DB.Delete(&computer, id)

	return computer
}
