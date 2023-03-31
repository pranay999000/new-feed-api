package models

type Config struct {
	Proxy		Proxy		`json:"proxy"`
	Databases	[]Database	`json:"databases"`
}