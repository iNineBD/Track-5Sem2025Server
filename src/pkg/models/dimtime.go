package models

type DimTime struct {
	IDTime    int64     `json:"id_time" gorm:"primaryKey;autoIncrement"`
	IDDay     int64     `json:"id_day"`
	DimDay    DimDay    `gorm:"foreignKey:IDDay;references:IDDay"`
	IDMonth   int64     `json:"id_month"`
	DimMonth  DimMonth  `gorm:"foreignKey:IDMonth;references:IDMonth"`
	IDYear    int64     `json:"id_year"`
	DimYear   DimYear   `gorm:"foreignKey:IDYear;references:IDYear"`
	IDHour    int64     `json:"id_hour"`
	DimHour   DimHour   `gorm:"foreignKey:IDHour;references:IDHour"`
	IDMinute  int64     `json:"id_minute"`
	DimMinute DimMinute `gorm:"foreignKey:IDMinute;references:IDMinute"`
}

func (DimTime) TableName() string {
	return "dim_time"
}
