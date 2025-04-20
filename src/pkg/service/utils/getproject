package utils

import (
	"inine-track/pkg/database"
	"inine-track/pkg/models"
)

func GetProject(idProject int64) (err error) {

	var project models.DimProject

	result := database.DB.Where("id = ?", idProject).First(&project)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
