package crud

import (
	"main/models"

	"gorm.io/gorm"
)

func DeletePeopleById(mydb *gorm.DB, personId int) (models.People, error) {
	var person models.People
	if err := mydb.Preload("Activity").First(&person, personId).Error; err != nil {
		return models.People{}, err
	}

	if err := mydb.Model(&person).Association("Activity").Delete(person.Activity); err != nil {
		return models.People{}, err
	}

	result := mydb.Delete(&person)
	if result.Error != nil {
		return models.People{}, result.Error
	}

	return person, nil
}

/////////////////////////////////////////////////////////////////////////////

func DeleteActivityById(mydb *gorm.DB, activityId int) (models.Activity, error) {
	var activity models.Activity
	if err := mydb.Preload("People").First(&activity, activityId).Error; err != nil {
		return models.Activity{}, err
	}

	if err := mydb.Model(&activity).Association("People").Delete(activity.People); err != nil {
		return models.Activity{}, err
	}

	result := mydb.Delete(&activity)
	if result.Error != nil {
		return models.Activity{}, result.Error
	}

	return activity, nil
}
