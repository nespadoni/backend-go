package lineup

import (
	"backend-go/internal/repository"
)

type Repository struct {
	repository *repository.BaseRepository
}

func NewLineUpRepository(repository *repository.BaseRepository) *Repository {
	return &Repository{repository: repository}
}
