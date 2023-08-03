package models

type Trainer struct {
	Id         uint16   `json:"id"`
	Name       string   `json:"name"`
	Region     string   `json:"region"`
	PokemonsId []uint16 `json:"pokemonsId"`
}
