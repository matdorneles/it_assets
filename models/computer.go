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
