package models

type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}
