package crud

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

func GetPeopleById(mydb *gorm.DB, personId int) (*models.People, error) {
	var person models.People
	err := mydb.Table("peoples").Where("id = ?", personId).Find(&person).Error
	if err != nil {
		return &models.People{}, fmt.Errorf("Person with ID %d not found: %v", personId, err)
	}
	return &person, nil
}

func GetPeopleByFullName(mydb *gorm.DB, personFullName string) (*models.People, error) {
	var person models.People
	err := mydb.Table("peoples").Where("full_name = ?", personFullName).First(&person).Error
	if err != nil {
		return &models.People{}, fmt.Errorf("Person with full name %s not found: %v", personFullName, err)
	}
	return &person, nil
}

func GetActivityById(mydb *gorm.DB, activityId int) (*models.Activity, error) {
	var activity models.Activity
	err := mydb.Table("activities").Where("id = ?", activityId).Find(&activity, activityId).Error
	if err != nil {
		return &models.Activity{}, fmt.Errorf("Activity with ID %d not found: %v", activityId, err)
	}
	return &activity, nil
}

func GetActivityByName(mydb *gorm.DB, activityName string) (*models.Activity, error) {
	var activity models.Activity
	err := mydb.Table("activities").Where("activity_name = ?", activityName).First(&activity).Error
	if err != nil {
		return &models.Activity{}, fmt.Errorf("Activity with name %s not found: %v", activityName, err)
	}
	return &activity, nil
}

func GetAllPeople(mydb *gorm.DB) ([]models.People, error) {
	var people []models.People
	err := mydb.Model(&models.People{}).Preload("Activity").Find(&people).Error
	return people, err
}

func GetAllActivities(mydb *gorm.DB) ([]models.Activity, error) {
	var activities []models.Activity
	err := mydb.Model(&models.Activity{}).Preload("People").Find(&activities).Error
	return activities, err
}