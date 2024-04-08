package postgres

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (p *Postgres) CreateMouse(ctx context.Context, mouse models.Mouse) error {
	err := p.Model(&models.Mouse{}).Create(mouse).Error
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}
	return nil
}

func (p *Postgres) DeleteMouse(ctx context.Context, id string) error {
	q := p.Model(&models.Mouse{}).
		Where("id = ?", id).
		Delete(models.Mouse{})
	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "mouse not found")
	}
	return nil
}

func (p *Postgres) GetMouse(ctx context.Context, mouse models.GetMouseReq) (models.Mouse, error) {
	var m models.Mouse
	q := p.Model(&models.Mouse{}).
		Where("id = ?", mouse.ID).
		Where("room_id = ?", mouse.RoomID).
		Where("building_id = ?", mouse.BuildingID)

	if mouse.Size != "" {
		q = q.Where("size = ?", mouse.Size)
	}
	if mouse.DPI != "" {
		q = q.Where("dpi = ?", mouse.DPI)
	}
	if mouse.Material != "" {
		q = q.Where("material = ?", mouse.Material)
	}
	if mouse.Sensor != "" {
		q = q.Where("sensor = ?", mouse.Sensor)
	}

	q.First(&m)
	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return models.Mouse{}, utils.NewError(utils.ErrNotFound, "mouse not found")
		}
		return models.Mouse{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}

	if q.Error != nil {
		return models.Mouse{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	return m, nil
}

func (p *Postgres) UpdateMouse(ctx context.Context, mouse models.Mouse) error {
	q := p.Model(&models.Mouse{}).
		Where("id = ?", mouse.ID).Updates(mouse)

	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "mouse not found")
	}
	return nil
}

func (p *Postgres) GetAllMouse(ctx context.Context, roomID, buildingID string) ([]models.Mouse, error) {
	var mice []models.Mouse

	q := p.Model(&models.Mouse{}).
		Where("room_id = ?", roomID).
		Where("building_id = ?", buildingID).
		Find(&mice)

	if q.Error != nil {
		return nil, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if len(mice) == 0 {
		return nil, utils.NewError(utils.ErrNotFound, "no mice found")
	}

	return mice, nil
}
