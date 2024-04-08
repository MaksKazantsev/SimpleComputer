package postgres

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (p *Postgres) CreateKeyboard(ctx context.Context, keyboard models.KeyBoard) error {
	err := p.Model(&models.KeyBoard{}).Create(keyboard).Error
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}
	return nil
}

func (p *Postgres) DeleteKeyboard(ctx context.Context, id string) error {
	q := p.Model(&models.KeyBoard{}).
		Where("id = ?", id).
		Delete(models.KeyBoard{})
	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "keyboard not found")
	}
	return nil
}

func (p *Postgres) UpdateKeyboard(ctx context.Context, keyboard models.KeyBoard) error {
	q := p.Model(&models.KeyBoard{}).
		Where("id = ?", keyboard.ID).
		Updates(keyboard)
	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "keyboard not found")
	}
	return nil
}

func (p *Postgres) GetKeyboard(ctx context.Context, keyboard models.GetKeyBoardReq) (models.KeyBoard, error) {
	var k models.KeyBoard
	q := p.Model(&models.KeyBoard{}).
		Where("id = ?", keyboard.ID).
		Where("room_id = ?", keyboard.RoomID).
		Where("building_id = ?", keyboard.BuildingID)

	if keyboard.Switches != "" {
		q = q.Where("switches = ?", keyboard.Switches)
	}
	if keyboard.Material != "" {
		q = q.Where("material = ?", keyboard.Material)
	}
	if keyboard.KeyType != "" {
		q = q.Where("key_type = ?", keyboard.KeyType)
	}
	if keyboard.Size != "" {
		q = q.Where("size = ?", keyboard.Size)
	}
	q.First(&k)

	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return models.KeyBoard{}, utils.NewError(utils.ErrNotFound, "keyboard not found")
		}
		return models.KeyBoard{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}

	if q.Error != nil {
		return models.KeyBoard{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	return k, nil
}

func (p *Postgres) GetAllKeyboard(ctx context.Context, roomID, buildingID string) ([]models.KeyBoard, error) {
	var keyboard []models.KeyBoard

	q := p.Model(&models.KeyBoard{}).
		Where("room_id = ?", roomID).
		Where("building_id = ?", buildingID).
		Find(&keyboard)
	if q.Error != nil {
		return nil, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if len(keyboard) == 0 {
		return nil, utils.NewError(utils.ErrNotFound, "no keyboards found")
	}

	return keyboard, nil
}
