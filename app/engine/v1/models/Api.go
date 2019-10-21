package models

import "time"

type Api struct {
	Id     string `json:"id"`
	Url     string `json:"url"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Active     int `json:"active"`
}
