package models

import "gorm.io/gorm"

type People struct {
	gorm.Model
	Id           int         `gorm:"primaryKey;autoIncrement"`
	Full_name    string      `gorm:"unique"`
	Phone_number string      `gorm:"unique"`
	Activity     []*Activity `gorm:"many2many:people_activities;"`
}

type Activity struct {
	gorm.Model
	Id           int `gorm:"primaryKey;autoIncrement"`
	ActivityName string
	Supervisor   string
	People       []*People `gorm:"many2many:people_activities;"`
}

type PeopleActivity struct {
	gorm.Model
	PeopleID   int
	ActivityID int
}
