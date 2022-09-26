package model

import "gorm.io/gorm"

type Mentee struct {
	gorm.Model
	Nama   string
	Alamat string
}

type MenteeModel struct {
	DB *gorm.DB
}
