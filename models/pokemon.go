package models

type Pokemon struct {
	Id    uint8  `json:"id"`
	Name  string `json:"name"`
	Hp    uint16 `json:"hp"`
	Def   uint16 `json:"def"`
	Defm  uint16 `json:"defm"`
	Atk   uint16 `json:"atk"`
	Spatk uint16 `json:"spatk"`
	Speed uint16 `json:"speed"`
}
