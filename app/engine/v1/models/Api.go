package models

type Api struct {
	Id     string `json:"id"`
	Url     string `json:"url"`
	Methods []string   `json:"methods"`
}
