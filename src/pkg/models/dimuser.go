package models

type DimUser struct {
	ID       int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName string  `json:"full_name"`
	Color    string  `json:"color"`
	Email    string  `json:"email"`
	IDRole   int64   `json:"ID_role"`
	DimRole  DimRole `gorm:"foreignKey;IdRole;references:Id"`
}

func (DimUser) TableName() string {
	return "dim_user"
}
