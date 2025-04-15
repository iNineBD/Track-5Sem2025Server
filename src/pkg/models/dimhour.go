package models

type DimHour struct {
	IDHour int64 `json:"id_hour" gorm:"primaryKey;autoIncrement"`
	Hour   int64 `json:"hour"`
}

func (DimHour) TableName() string {
	return "dim_hour"
}
