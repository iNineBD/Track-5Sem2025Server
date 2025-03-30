package models

type DimTag struct {
	Id         int64  `json:"id" gorm:"primryKey;autoIncrement"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Id_card    int64  `json:"id_card"`
	Id_project int64  `json:"id_project"`
}

func (DimTag) TableName() string {
	return "dim_tag"
}
