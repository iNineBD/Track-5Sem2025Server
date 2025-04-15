package models

type DimTag struct {
	IDTag   int64  `json:"id_tag" gorm:"primryKey;autoIncrement"`
	NameTag string `json:"name_tag"`
}

func (DimTag) TableName() string {
	return "dim_tag"
}
