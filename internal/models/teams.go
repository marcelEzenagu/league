package models

type Team struct {
	Name      string   `json:"name" form:"name"`
	Logo      string   `json:"logo" form:"logo"`
	Position  string   `json:"position" form:"position"`
	Played    string   `json:"played" form:"played"`
	Won       string   `json:"won" form:"won"`
	Lost      string   `json:"lost" form:"lost"`
	Drawn     string   `json:"drawn" form:"drawn"`
	Players   []Player `json:"players,omitempty"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	DeletedAt string   `json:"deletedAt"`
}

type Player struct {
	FirstName      string `json:"firstName" form:"firstName"`
	LastName       string `json:"lastName" form:"lastName"`
	ProfilePicture string `json:"profilePicture" form:"profilePicture"`
	Email          string `json:"email" form:"email"`
	Position       string `json:"position"  form:"position"`
}
