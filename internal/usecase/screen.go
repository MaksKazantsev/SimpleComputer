package usecase

import (
	"college/internal/usecase/models"
	"context"
	"fmt"
	"time"
)

func (u *UseCase) CreateScreen(ctx context.Context, screen models.Screen) error {
	screen.CreatedAt = time.Now()
	screen.UpdatedAt = time.Now()

	if err := u.repo.CreateScreen(ctx, screen); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) DeleteScreen(ctx context.Context, id string) error {
	if err := u.repo.DeleteScreen(ctx, id); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) UpdateScreen(ctx context.Context, screen models.Screen) error {
	screen.UpdatedAt = time.Now()
	if err := u.repo.UpdateScreen(ctx, screen); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) GetScreen(ctx context.Context, screen models.GetScreenReq) (models.Screen, error) {
	sc, err := u.repo.GetScreen(ctx, screen)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return sc, nil
}

func (u *UseCase) GetAllScreen(ctx context.Context, roomID, buildingID string) ([]models.Screen, error) {
	screens, err := u.repo.GetAllScreen(ctx, roomID, buildingID)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return screens, nil
}
