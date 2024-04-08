package usecase

import (
	"college/internal/usecase/models"
	"context"
	"fmt"
	"time"
)

func (u *UseCase) CreateKeyboard(ctx context.Context, keyboard models.KeyBoard) error {
	keyboard.CreatedAt = time.Now()
	keyboard.UpdatedAt = time.Now()

	if err := u.repo.CreateKeyboard(ctx, keyboard); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) DeleteKeyBoard(ctx context.Context, id string) error {
	if err := u.repo.DeleteKeyboard(ctx, id); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) UpdateKeyBoard(ctx context.Context, keyboard models.KeyBoard) error {
	keyboard.UpdatedAt = time.Now()
	if err := u.repo.UpdateKeyboard(ctx, keyboard); err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return nil
}

func (u *UseCase) GetKeyboard(ctx context.Context, keyboard models.GetKeyBoardReq) (models.KeyBoard, error) {
	key, err := u.repo.GetKeyboard(ctx, keyboard)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return key, nil
}

func (u *UseCase) GetAllKeyBoard(ctx context.Context, roomID, buildingID string) ([]models.KeyBoard, error) {
	keyboard, err := u.repo.GetAllKeyboard(ctx, roomID, buildingID)
	if err != nil {
		_ = fmt.Errorf("repo error: %v", err)
	}
	return keyboard, nil
}
