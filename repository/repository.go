package repository

import "github.com/viniosilva/go-api/repository/cat"

type Repository interface {
	Cat() cat.CatRepository
}

type repository struct {
	cat cat.CatRepository
}

func NewRepository(cat cat.CatRepository) Repository {
	return &repository{cat: cat}
}

func (repo *repository) Cat() cat.CatRepository { return repo.cat }
