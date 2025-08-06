package models

type Tasks struct {
	Id          int    `json:"id"` // ! struct tags
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
