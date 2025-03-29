package models

type DimStatus struct {
	ID   int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}

func (DimStatus) TableName() string {
	return "dim_status"
}
