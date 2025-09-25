package entities

import "time"

type Team struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LogoURL     string `json:"logoURL"`
	Country     string `json:"country"`
	Tag         string `json:"tag"`
	Description string `json:"description"`

	Players []Player `json:"players"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
