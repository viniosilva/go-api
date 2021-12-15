package cat

import (
	"time"

	"github.com/viniosilva/go-api/model"
	repo "github.com/viniosilva/go-api/repository"
)

//go:generate mockgen -destination=cat_app_mock.go -package=cat . CatApp
type CatApp interface {
	FindCats() ListCatsDto
	CreateCat(cmd CatCmd) (CatDto, error)
}

type catAppImpl struct {
	repo repo.CatRepository
}

func NewCatApp(catRepository repo.CatRepository) CatApp {
	return &catAppImpl{repo: catRepository}
}

func (impl catAppImpl) CreateCat(cmd CatCmd) (CatDto, error) {
	catBirthday, err := time.Parse("2006-01-02", cmd.Birthday)

	if err != nil {
		return CatDto{}, &InvalidCmdError{}
	}

	cat := impl.repo.Create(&model.Cat{
		Birthday: catBirthday,
		Name:     cmd.Name,
	})

	return CatDto{
		ID:       cat.ID,
		Birthday: cat.Birthday.Format("2006-01-02"),
		Name:     cat.Name,
	}, nil
}

func (impl catAppImpl) FindCats() ListCatsDto {
	cats := impl.repo.Find()

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
