package service

import (
	"context"
	"errors"
	"ramazon/constants"
	"ramazon/models"
	"ramazon/pkg/logger"
	"ramazon/storage"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ramazonServiceImpl struct {
	storage storage.IStorage
	logger  logger.Logger
}

func NewRamazonService(log logger.Logger) RamazonService {
	return &ramazonServiceImpl{
		storage: storage.NewStoragePg(),
		logger:  log,
	}
}

type RamazonService interface {
	SetPrayTime(ctx context.Context, req models.Ramazon) error
}

func (s *ramazonServiceImpl) SetPrayTime(ctx context.Context, req models.Ramazon) error {
	err := s.storage.Ramazon().SetPrayTime(ctx, req)
	if err != nil {
		s.logger.Error("error in Set Pray Time: ", zap.Error(err))
		if errors.Is(err, constants.ErrRowsAffectedIsZero) {
			s.logger.Error("error rows affected is not zero")
			return status.Error(codes.InvalidArgument, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
