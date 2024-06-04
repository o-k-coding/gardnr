package grdnr

import (
	"context"

	"okcoding.com/grdnr/internal/config"
	"okcoding.com/grdnr/internal/objectstorage"
)

type Service struct {
	Config  config.GrdnrConfig
	Storage objectstorage.ObjectStorage
}

var (
	Grdnr Service
)

func (s *Service) Init(ctx context.Context) (err error) {
	s.Config, err = config.InitConfig()
	if err != nil {
		return err
	}
	s.Storage, err = objectstorage.NewObjectStorage(ctx, s.Config)
	if err != nil {
		return err
	}
	return nil
}
