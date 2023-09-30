package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	awsEndpoint string
	bucketName  string

	s3svc *s3.Client
)

func init() {
	awsEndpoint = os.Getenv("AWS_ENDPOINT")
	bucketName = os.Getenv("S3_BUCKET")

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:       "aws",
				URL:               awsEndpoint,
				HostnameImmutable: true,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	s3svc = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}

func main() {
	if _, err := s3svc.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}); err != nil {
		log.Fatalf("Cannot create bucket: %s", err)
	}
}
