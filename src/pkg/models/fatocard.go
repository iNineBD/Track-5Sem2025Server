package models

type FatoCard struct {
	IDStatus   int64      `json:"id_status"`
	DimStatus  DimStatus  `gorm:"foreignKey;IdStatus;references:Id"`
	IDTag      int64      `json:"id_tag"`
	DimTag     DimTag     `gorm:"foreignKey;IdTag;references:Id"`
	IDUser     int64      `json:"id_user"`
	DimUser    DimUser    `gorm:"foreignKey;IdUser;references:Id"`
	IDProject  int64      `json:"id_project"`
	DimProject DimProject `gorm:"foreignKey;IdProject;references:Id"`
	QtdCard    int64      `json:"qtd_card"`
}
