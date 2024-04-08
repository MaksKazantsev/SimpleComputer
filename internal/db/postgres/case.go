package postgres

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (p *Postgres) CreateCase(ctx context.Context, pc models.Case) error {
	err := p.Model(&models.Case{}).Create(pc).Error
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}
	return nil
}

func (p *Postgres) DeleteCase(ctx context.Context, id string) error {
	q := p.Model(&models.Case{}).Where("id = ?", id).Delete(models.Case{})
	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "case not found")
	}
	return nil
}

func (p *Postgres) UpdateCase(ctx context.Context, pc models.Case) error {
	q := p.Model(&models.Case{}).
		Where("id = ?", pc.ID).
		Updates(pc)
	if q.Error != nil {
		return utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if q.RowsAffected == 0 {
		return utils.NewError(utils.ErrNotFound, "case not found")
	}
	return nil
}

func (p *Postgres) GetCase(ctx context.Context, pc models.GetCaseReq) (models.Case, error) {
	var c models.Case
	q := p.Model(&models.Case{}).
		Where("id = ?", pc.ID).
		Where("room_id = ?", pc.RoomID).
		Where("building_id = ?", pc.BuildingID)

	if pc.Size != "" {
		q = q.Where("size = ?", pc.Size)
	}
	if pc.Color != "" {
		q = q.Where("color = ?", pc.Color)
	}
	if pc.Compatibility != "" {
		q = q.Where("compatibility = ?", pc.Compatibility)
	}
	q.First(&c)
	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return models.Case{}, utils.NewError(utils.ErrNotFound, "case not found")
		}
		return models.Case{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}

	if q.Error != nil {
		return models.Case{}, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	return c, nil
}

func (p *Postgres) GetAllCase(ctx context.Context, roomID, buildingID string) ([]models.Case, error) {
	var cases []models.Case

	q := p.Model(&models.Case{}).
		Where("room_id = ?", roomID).
		Where("building_id = ?", buildingID).
		Find(&cases)

	if q.Error != nil {
		return nil, utils.NewError(utils.ErrInternal, q.Error.Error())
	}
	if len(cases) == 0 {
		return nil, utils.NewError(utils.ErrNotFound, "no cases found")
	}

	return cases, nil
}
