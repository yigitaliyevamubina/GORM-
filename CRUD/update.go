package crud

import (
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

func UpdatePersonFullNameByid(mydb *gorm.DB, personId int, newName string) error {
	var existingPerson models.People
	err := mydb.First(&existingPerson, personId).Error
	if err != nil {
		return fmt.Errorf("Person with ID %d not found: %s", personId, err)
	}

	result := mydb.Model(&existingPerson).Where("id = ?", personId).Update("full_name", newName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdatePersonPhoneNumberById(mydb *gorm.DB, personId int, newPhoneNumber string) error {
	var existingPerson models.People
	err := mydb.First(&existingPerson, personId).Error
	if err != nil {
		return fmt.Errorf("Person with ID %d not found: %s", personId, err)
	}

	result := mydb.Model(&existingPerson).Where("id = ?", personId).Update("phone_number", newPhoneNumber)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateActivityNameById(mydb *gorm.DB, activityId int, newActivityName string) error {
	var existingActivity models.Activity
	err := mydb.First(&existingActivity, activityId)
	if err != nil {
		return fmt.Errorf("Activity with ID %d not found: %v", activityId, err)
	}

	result := mydb.Model(&existingActivity).Where("id = ?", activityId).Update("activity_name", newActivityName)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("No activity updated with ID %d", activityId)
	}

	return nil
}

func UpdateActivitySupervisorById(mydb *gorm.DB, activityId int, newSupervisor string) error {
	var existingActivity models.Activity
	err := mydb.First(&existingActivity, activityId)
	if err != nil {
		return fmt.Errorf("Activity with ID %d not found: %v", activityId, err)
	}

	result := mydb.Model(&existingActivity).Where("id = ?", activityId).Update("supervisor", newSupervisor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
