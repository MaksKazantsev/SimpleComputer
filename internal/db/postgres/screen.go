package postgres

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (p *Postgres) CreateScreen(ctx context.Context, screen models.Screen) error {
	err := p.Model(&models.Screen{}).Create(screen).Error
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}
	return nil
}

func (p *Postgres) DeleteScreen(ctx context.Context, id string) error {
	q := p.Model(&models.Screen{}).
		Where("id = ?", id).
		Delete(models.Screen{})
	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "screen not found")
	}
	return nil
}

func (p *Postgres) UpdateScreen(ctx context.Context, screen models.Screen) error {
	q := p.Model(&models.Screen{}).
		Where("id = ?", screen.ID).
		Updates(screen)
	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "screen not found")
	}
	return nil
}

func (p *Postgres) GetScreen(ctx context.Context, screen models.GetScreenReq) (models.Screen, error) {
	var s models.Screen
	q := p.Model(&models.Screen{}).
		Where("id = ?", screen.ID).
		Where("room_id = ?", screen.RoomID).
		Where("building_id = ?", screen.BuildingID)

	if screen.Size != "" {
		q = q.Where("size = ?", screen.Size)
	}
	if screen.Herz != "" {
		q = q.Where("herz = ?", screen.Herz)
	}
	if screen.Matrix != "" {
		q = q.Where("matrix = ?", screen.Matrix)
	}

	q.First(&s)
	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return models.Screen{}, utils.NewError(utils.ErrNotFound, "screen not found")
		}
		return models.Screen{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}

	if q.Error != nil {
		return models.Screen{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	return s, nil
}

func (p *Postgres) GetAllScreen(ctx context.Context, roomID, buildingID string) ([]models.Screen, error) {
	var screens []models.Screen

	q := p.Model(&models.Screen{}).
		Where("room_id = ?", roomID).
		Where("building_id = ?", buildingID).
		Find(&screens)
	if q.Error != nil {
		return nil, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if len(screens) == 0 {
		return nil, utils.NewError(utils.ErrNotFound, "no screens found")
	}

	return screens, nil
}
