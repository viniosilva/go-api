package cat

type CatDto struct {
	ID       int    `json:"id"`
	Birthday string `json:"birthday"`
	Name     string `json:"name"`
}

type ListCatsDto struct {
	Data []CatDto `json:"data"`
}

type CatCmd struct {
	Birthday string `json:"birthday" example:"2000-01-01"`
	Name     string `json:"name" example:"Mimoso"`
}
