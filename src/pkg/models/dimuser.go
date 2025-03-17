package models

type DimUser struct {
	Id       int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName string  `json:"full_name"`
	Color    string  `json:"color"`
	Email    string  `json:"email"`
	IdRole   int64   `json:"id_role"`
	DimRole  DimRole `gorm:"foreignKey;IdRole;references:Id"`
}

func (DimUser) TableName() string {
	return "dim_user"
}
