package models

type DimStatus struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name"`
	IdCard    int64  `json:"id_card"`
	IdProject int64  `json:"id_project"`
}

func (DimStatus) TableName() string {
	return "dim_status"
}
