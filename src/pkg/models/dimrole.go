package models

type DimRole struct {
	IDRole   int64  `json:"id_role" gorm:"primaryKey;autoIncrement"`
	NameRole string `json:"name_role"`
}

func (DimRole) TableName() string {
	return "dim_role"
}
