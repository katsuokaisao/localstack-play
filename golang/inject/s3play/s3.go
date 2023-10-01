package s3play

import (
	"fmt"

	"github.com/google/wire"
	"github.com/katsuokaisao/localstack/infra/s3"
)

var s3Set = wire.NewSet(
	s3.NewS3Repository,
	provideS3Component,
)

func provideS3Component(cfg *Config) (*s3.S3Component, error) {
	if cfg.AWS == nil {
		return nil, fmt.Errorf("AWS config is nil")
	}

	return &s3.S3Component{
		Endpoint: cfg.AWS.Endpoint,
	}, nil
}
