package models

import "time"

type DimProject struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Description  string    `json:"description"`
	Name         string    `json:"name"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
}

func (DimProject) TableName() string {
	return "dim_project"
}
