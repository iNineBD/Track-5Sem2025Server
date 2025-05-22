package models

type DimPlatform struct {
	IDPlatform   int64  `json:"id_platform" gorm:"primaryKey;autoIncrement"`
	NamePlatform string `json:"name_platform"`
}

func (DimPlatform) TableName() string {
	return "dim_platform"
}
