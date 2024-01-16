package crud

import (
	"main/models"

	"gorm.io/gorm"
)

func CreatePeople(mydb *gorm.DB, people []*models.People) error {
	result := mydb.Create(&people)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateActivity(mydb *gorm.DB, activities []*models.Activity) error {
	result := mydb.Create(&activities)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
