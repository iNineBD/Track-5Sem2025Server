package models

type DimRole struct {
	ID   int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}

func (DimRole) TableName() string {
	return "dim_role"
}
