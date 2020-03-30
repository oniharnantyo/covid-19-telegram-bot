package models

type ErrorModel struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}
