package models

type UserAuth struct {
	Username string `json:"login"`
	Password string `json:"password,omitempty"`
}

type Game struct {
	Title string `json:"title,omitempty"`
	Genre string `json:"genre,omitempty"`
}
