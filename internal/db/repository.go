package db

import (
	"college/internal/usecase/models"
	"context"
)

type Repository interface {
	CreateMouse(ctx context.Context, mouse models.Mouse) error
	DeleteMouse(ctx context.Context, id string) error
	UpdateMouse(ctx context.Context, mouse models.Mouse) error
	GetMouse(ctx context.Context, mouse models.GetMouseReq) (models.Mouse, error)
	GetAllMouse(ctx context.Context, roomID, buildingID string) ([]models.Mouse, error)

	CreateKeyboard(ctx context.Context, keyboard models.KeyBoard) error
	DeleteKeyboard(ctx context.Context, id string) error
	UpdateKeyboard(ctx context.Context, keyboard models.KeyBoard) error
	GetKeyboard(ctx context.Context, keyboard models.GetKeyBoardReq) (models.KeyBoard, error)
	GetAllKeyboard(ctx context.Context, roomID, buildingID string) ([]models.KeyBoard, error)

	CreateScreen(ctx context.Context, screen models.Screen) error
	DeleteScreen(ctx context.Context, id string) error
	UpdateScreen(ctx context.Context, screen models.Screen) error
	GetScreen(ctx context.Context, screen models.GetScreenReq) (models.Screen, error)
	GetAllScreen(ctx context.Context, roomID, buildingID string) ([]models.Screen, error)

	CreateCase(ctx context.Context, pc models.Case) error
	DeleteCase(ctx context.Context, id string) error
	UpdateCase(ctx context.Context, pc models.Case) error
	GetCase(ctx context.Context, pc models.GetCaseReq) (models.Case, error)
	GetAllCase(ctx context.Context, roomID, buildingID string) ([]models.Case, error)
}
