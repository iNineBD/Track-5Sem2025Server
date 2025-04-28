package statisticsdto

type UserData struct {
	NameUser string `json:"name_user" gorm:"column:name_user"`
	Qtd      int64  `json:"qtd" gorm:"column:qtd_card_user"`
}
