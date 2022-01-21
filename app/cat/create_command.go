package cat

import (
	"time"

	"github.com/viniosilva/go-api/exception"
	"github.com/viniosilva/go-api/model"
)

type CreateDto struct {
	Birthday string
	Name     string
}

type CreateResult struct {
	ID       int
	Birthday time.Time
	Name     string
}

func (app *catApp) Create(dto CreateDto) (CreateResult, error) {
	catBirthday, err := time.Parse("2006-01-02", dto.Birthday)
	if err != nil {
		return CreateResult{}, &exception.InvalidDateException{}
	}

	catModel := app.repo.Create(&model.Cat{
		Birthday: catBirthday,
		Name:     dto.Name,
	})

	return CreateResult{
		ID:       catModel.ID,
		Birthday: catModel.Birthday,
		Name:     catModel.Name,
	}, nil
}
