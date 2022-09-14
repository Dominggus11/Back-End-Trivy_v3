package models

type Dockerfiles struct {
	ID       uint   `json:"id"`
	Pathfile string `json:"pathfile"`
	PathJson string `json:"pathjson"`
}
