package models

type DimDay struct {
	IDDay int64 `json:"id_day" gorm:"primaryKey;autoIncrement"`
	Day   int64 `json:"day"`
}

func (DimDay) TableName() string {
	return "dim_day"
}
