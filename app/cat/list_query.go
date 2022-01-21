package cat

import (
	"time"
)

type CatResult struct {
	ID       int
	Birthday time.Time
	Name     string
}

type ListResult struct {
	Data []CatResult
}

func (app *catApp) List() (ListResult, error) {
	cats := app.repo.Find()

	data := []CatResult{}
	for _, cat := range cats {
		data = append(data, CatResult{
			ID:       cat.ID,
			Name:     cat.Name,
			Birthday: cat.Birthday,
		})
	}

	return ListResult{Data: data}, nil
}
