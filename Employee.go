package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name string `json:"name"`
	City string `json:"city"`
}
