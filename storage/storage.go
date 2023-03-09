package storage

import (
	"ramazon/storage/repo"

	"ramazon/storage/postgres"
)

type IStorage interface {
	Ramazon() repo.RamazonI
}

type storagePg struct {
	// db           *sqlx.DB
	ramazonRepo repo.RamazonI
}

func NewStoragePg() IStorage {
	return &storagePg{
		// db:           db,
		ramazonRepo: postgres.NewRamazonRepo(),
	}
}

func (s storagePg) Ramazon() repo.RamazonI {
	return s.ramazonRepo
}
