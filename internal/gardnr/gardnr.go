package gardnr

import (
	"context"

	"okcoding.com/gardnr/internal/config"
	"okcoding.com/gardnr/internal/objectstorage"
)

type Service struct {
	Config  config.GardnrConfig
	Storage objectstorage.ObjectStorage
}

var (
	Gardnr Service
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
