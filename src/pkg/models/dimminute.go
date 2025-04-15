package models

type DimMinute struct {
	IDMinute int64 `json:"id_minute" gorm:"primaryKey;autoIncrement"`
	Minute   int64 `json:"minute"`
}

func (DimMinute) TableName() string {
	return "dim_minute"
}
