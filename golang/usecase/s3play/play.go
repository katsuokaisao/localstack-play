package s3play

import (
	"github.com/katsuokaisao/localstack/domain/repository"
	"github.com/rs/zerolog/log"
)

type S3PlayUseCase interface {
	Play()
}

type s3PlayUseCase struct {
	S3Repository repository.S3Repository
}

func NewS3PlayUseCase(
	s3Repository repository.S3Repository,
) S3PlayUseCase {
	return &s3PlayUseCase{
		S3Repository: s3Repository,
	}
}

func (u *s3PlayUseCase) Play() {
	if err := u.S3Repository.CreateBucket("test"); err != nil {
		log.Err(err)
	}
}
