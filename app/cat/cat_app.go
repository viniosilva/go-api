package cat

import (
	repo "github.com/viniosilva/go-api/repository"
)

//go:generate mockgen -destination=cat_app_mock.go -package=app . CatApp
type CatApp interface {
	FindCats() ListCatsDto
}

type CatAppImpl struct {
	repo repo.CatRepository
}

func NewCatApp(catRepository repo.CatRepository) CatApp {
	return &CatAppImpl{repo: catRepository}
}

func (impl CatAppImpl) FindCats() ListCatsDto {
	cats := impl.repo.FindCats()

	data := []CatDto{}
	for _, cat := range cats {
		data = append(data, CatDto{
			ID:       cat.ID,
			Name:     cat.Name,
			Birthday: cat.Birthday.Format("2006-01-02"),
		})
	}

	return ListCatsDto{Data: data}
}
