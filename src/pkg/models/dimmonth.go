package models

type DimMonth struct {
	IDMonth int64 `json:"id_month" gorm:"primaryKey;autoIncrement"`
	Month   int64 `json:"month"`
}

func (DimMonth) TableName() string {
	return "dim_month"
}
