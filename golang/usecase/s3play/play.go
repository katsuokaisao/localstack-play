package s3play

import (
	"log"

	"github.com/katsuokaisao/localstack/domain/repository"
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
	var (
		exists bool
		err    error
	)

	exists, err = u.S3Repository.BucketExists("test")
	if err != nil {
		log.Printf("failed to check bucket existence: %v\n", err)
	}
	log.Printf("bucket exists: %v\n", exists)

	if !exists {
		if err := u.S3Repository.CreateBucket("test"); err != nil {
			log.Printf("failed to create bucket: %v\n", err)
		} else {
			log.Println("bucket created")
		}
	}
}
