package models

type Trainer struct {
	Id       uint16    `json:"id"`
	Name     string    `json:"name"`
	Region   string    `json:"region"`
	Pokemons []Pokemon `json:"pokemons"`
	Team     []Team    `json:"team"`
}
