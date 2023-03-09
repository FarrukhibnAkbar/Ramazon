package repo

import (
	"context"

	"ramazon/models"
)

type RamazonI interface {
	SetPrayTime(ctx context.Context, model models.Ramazon) error
}
