package statisticsdto

type TagData struct {
	Tag string `json:"tag_name" gorm:"column:name_tag"`
	Qtd int64  `json:"qtd" gorm:"column:qtd_card_tag"`
}
