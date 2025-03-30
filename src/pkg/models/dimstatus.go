package models

type DimStatus struct {
	Id         int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	Id_card    int64  `json:"id_card"`
	Id_project int64  `json:"id_project"`
}

func (DimStatus) TableName() string {
	return "dim_status"
}
