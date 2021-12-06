package cat

type CatDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

type ListCatsDto struct {
	Data []CatDto `json:"data"`
}
