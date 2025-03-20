package projectdto

import "time"

type GetProjectsResponse struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
	FinishDate   time.Time `json:"finish_date"`
	LogoBigUrl   string    `json:"logo_big_url"`
	LogoSmallUrl string    `json:"logo_small_url"`
}
