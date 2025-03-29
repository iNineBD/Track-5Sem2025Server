package models

type DimTag struct {
	Id    int64  `json:"id" gorm:"primryKey;autoIncrement"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (DimTag) TableName() string {
	return "dim_tag"
}
