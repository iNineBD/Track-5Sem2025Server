package models

type DimStatus struct {
	IDStatus   int64  `json:"id_status" gorm:"primaryKey;autoIncrement"`
	NameStatus string `json:"name_status"`
}

func (DimStatus) TableName() string {
	return "dim_status"
}
