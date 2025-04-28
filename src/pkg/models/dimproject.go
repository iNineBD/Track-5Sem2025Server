package models

type DimProject struct {
	IDProject   int64  `json:"id_project" gorm:"primaryKey;autoIncrement"`
	NameProject string `json:"name_project"`
	Description string `json:"description"`
}

func (DimProject) TableName() string {
	return "dim_project"
}
