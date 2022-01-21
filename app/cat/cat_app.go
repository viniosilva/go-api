package cat

import "github.com/viniosilva/go-api/repository/cat"

//go:generate mockgen -destination=cat_app_mock.go -package=cat . CatApp
type CatApp interface {
	Create(dto CreateDto) (CreateResult, error)
	List() (ListResult, error)
}

type catApp struct {
	repo cat.CatRepository
}

func NewApp(repo cat.CatRepository) CatApp {
	return &catApp{repo: repo}
}
