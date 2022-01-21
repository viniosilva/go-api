package cat

import (
	"github.com/viniosilva/go-api/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=cat_repository_mock.go -package=cat . CatRepository
type CatRepository interface {
	Create(cat *model.Cat) *model.Cat
	Find() []model.Cat
}

type catRepository struct {
	gormDB *gorm.DB
}

func NewRepository(gormDB *gorm.DB) CatRepository {
	return &catRepository{gormDB: gormDB}
}
