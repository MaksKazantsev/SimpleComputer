package usecase

import (
	"college/internal/db"
)

type UseCase struct {
	repo db.Repository
}

func NewUseCase(repo db.Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
