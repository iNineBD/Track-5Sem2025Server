package models

type FatoCard struct {
	IDFatoCard     int64      `json:"id_fato_card" gorm:"primaryKey;autoIncrement"`
	IDCard         int64      `json:"id_card"`
	DimCard        DimCard    `gorm:"foreignKey;IDCard;references:IDCard"`
	IDProject      int64      `json:"id_project"`
	DimProject     DimProject `gorm:"foreignKey;IDProject;references:IDProject"`
	IDUser         int64      `json:"id_user"`
	DimUser        DimUser    `gorm:"foreignKey;IDUser;references:IDUser"`
	IDStatus       int64      `json:"id_status"`
	DimStatus      DimStatus  `gorm:"foreignKey;IDStatus;references:IDStatus"`
	IDTimeCreated  int64      `json:"id_time_created"`
	DimTimeCreated DimTime    `gorm:"foreignKey;IDTimeCreated;references:IDTime"`
	IDTag          int64      `json:"id_tag"`
	DimTag         DimTag     `gorm:"foreignKey;IDTag;references:IDTag"`
	Qtdcards       int64      `json:"qtd_card"`
}

func (FatoCard) TableName() string {
	return "fato_cards"
}
