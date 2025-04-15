package models

type DimCard struct {
	IDCard      int64  `json:"id_card" gorm:"primaryKey;autoIncrement"`
	NameCard    string `json:"name_card"`
	Description string `json:"description"`
}

func (DimCard) TableName() string {
	return "dim_card"
}
