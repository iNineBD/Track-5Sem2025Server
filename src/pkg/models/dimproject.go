package models

import "time"

type DimProject struct {
	Id           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Description  string    `json:"description"`
	Name         string    `json:"name"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
	FinishDate   time.Time `json:"finish_date"`
	LogoBigUrl   string    `json:"logo_big_url"`
	LogoSmallUrl string    `json:"logo_small_url"`
}

func (DimProject) TableName() string {
	return "dim_project"
}
