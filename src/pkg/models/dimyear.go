package models

type DimYear struct {
	IDYear int64 `json:"id_year"`
	Year   int64 `json:"year"`
}

func (DimYear) TableName() string {
	return "dim_year"
}
