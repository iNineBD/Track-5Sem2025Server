package models

type DimProject struct {
	IDProject   int64       `json:"id_project" gorm:"primaryKey;autoIncrement"`
	NameProject string      `json:"name_project"`
	Description string      `json:"description"`
	IDPlatform  int64       `json:"id_platform"`
	DimPlatform DimPlatform `gorm:"foreignKey:IDPlatform;references:IDPlatform"`
}

func (DimProject) TableName() string {
	return "dim_project"
}
