package model

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Nama  string
	Hobby string
}

type MentorModel struct {
	DB *gorm.DB
}

func GetAll() {

}

func Search() {

}
