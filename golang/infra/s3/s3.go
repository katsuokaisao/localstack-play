package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/katsuokaisao/localstack/domain/repository"
)

type s3Repository struct {
	Client *s3.Client
}

type S3Component struct {
	Endpoint string
}

func NewS3Repository(s3Component *S3Component) (repository.S3Repository, error) {
	client, err := newS3Client(s3Component)
	if err != nil {
		return nil, fmt.Errorf("cannot create s3 client: %s", err)
	}

	return &s3Repository{
		Client: client,
	}, nil
}

func newS3Client(s3Component *S3Component) (*s3.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if s3Component.Endpoint != "" {
			return aws.Endpoint{
				PartitionID: "aws",
				URL:         s3Component.Endpoint,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot load the AWS configs: %s", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return client, nil
}
