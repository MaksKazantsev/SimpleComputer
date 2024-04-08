package usecase

import (
	"college/internal/usecase/models"
	"context"
	"fmt"
	"time"
)

func (u *UseCase) CreateMouse(ctx context.Context, mouse models.Mouse) error {
	mouse.CreatedAt = time.Now()
	mouse.UpdatedAt = time.Now()

	if err := u.repo.CreateMouse(ctx, mouse); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) DeleteMouse(ctx context.Context, id string) error {
	if err := u.repo.DeleteMouse(ctx, id); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) UpdateMouse(ctx context.Context, mouse models.Mouse) error {
	mouse.UpdatedAt = time.Now()
	if err := u.repo.UpdateMouse(ctx, mouse); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) GetMouse(ctx context.Context, mouse models.GetMouseReq) (models.Mouse, error) {
	mo, err := u.repo.GetMouse(ctx, mouse)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return mo, nil
}

func (u *UseCase) GetAllMouse(ctx context.Context, roomID, buildingID string) ([]models.Mouse, error) {
	mice, err := u.repo.GetAllMouse(ctx, roomID, buildingID)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return mice, nil
}
