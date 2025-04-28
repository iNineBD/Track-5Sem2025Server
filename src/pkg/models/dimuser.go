package models

type DimUser struct {
	IDUser   int64   `json:"id_user" gorm:"primaryKey;autoIncrement"`
	NameUser string  `json:"name_user"`
	Email    string  `json:"email"`
	IDRole   int64   `json:"id_role"`
	DimRole  DimRole `gorm:"foreignKey:IDRole;references:IDRole"`
	Password string  `json:"password"`
}

func (DimUser) TableName() string {
	return "dim_user"
}
