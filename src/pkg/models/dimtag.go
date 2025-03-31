package models

type DimTag struct {
	ID        int64  `json:"id" gorm:"primryKey;autoIncrement"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	IDCard    int64  `json:"id_card"`
	IDProject int64  `json:"id_project"`
}

func (DimTag) TableName() string {
	return "dim_tag"
}
