package usecase

import (
	"college/internal/usecase/models"
	"context"
	"fmt"
	"time"
)

func (u *UseCase) CreateCase(ctx context.Context, pc models.Case) error {
	pc.CreatedAt = time.Now()
	pc.UpdatedAt = time.Now()

	if err := u.repo.CreateCase(ctx, pc); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}

	return nil
}

func (u *UseCase) DeleteCase(ctx context.Context, id string) error {
	if err := u.repo.DeleteCase(ctx, id); err != nil {
		_ = fmt.Errorf("repo err: %v", err)
	}
	return nil
}

func (u *UseCase) UpdateCase(ctx context.Context, pc models.Case) error {
	pc.UpdatedAt = time.Now()

	if err := u.repo.UpdateCase(ctx, pc); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}

	return nil
}

func (u *UseCase) GetCase(ctx context.Context, pc models.GetCaseReq) (models.Case, error) {
	c, err := u.repo.GetCase(ctx, pc)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}

	return c, nil
}

func (u *UseCase) GetAllCase(ctx context.Context, roomID, buildingID string) ([]models.Case, error) {
	cases, err := u.repo.GetAllCase(ctx, roomID, buildingID)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return cases, nil
}
