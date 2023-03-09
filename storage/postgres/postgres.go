package postgres

import (
	"context"
	"ramazon/constants"
	"ramazon/models"
	"ramazon/platforma/postgres"
	"ramazon/storage/repo"

	"gorm.io/gorm"
)

type ramazonRepo struct {
	db *gorm.DB
}

func NewRamazonRepo() repo.RamazonI {
	return &ramazonRepo{
		db: postgres.DB(),
	}
}

func (r *ramazonRepo) SetPrayTime(ctx context.Context, model models.Ramazon) error {
	res := r.db.Table("pray_time").Create(&model)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected == 0 {
		return constants.ErrRowsAffectedIsZero
	}

	return nil
}
