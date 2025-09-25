package entities

import "time"

type Player struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Nickname    string `json:"nickname"`
	Nationality string `json:"nationality"`
	PhotoURL    string `json:"photoURL"`
	BaseRating  int32  `json:"baseRating"`
	Role        string `json:"role"`

	// Dynamic Attributes
	FatigueBps   int32 `json:"fatigueBps"`
	MoraleBps    int32 `json:"moraleBps"`
	HappinessBps int32 `json:"happinessBps"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
