package s3

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

func (r *s3Repository) CreateBucket(bucketName string) error {
	_, err := r.Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintApNortheast1,
		},
	})
	if err != nil {
		return fmt.Errorf(
			"failed to create bucket %v in Region %v. Here's why: %v",
			bucketName, types.BucketLocationConstraintApNortheast1, err,
		)
	}

	return nil
}

func (r *s3Repository) BucketExists(bucketName string) (bool, error) {
	_, err := r.Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err == nil {
		return true, nil
	}

	var apiError smithy.APIError
	if errors.As(err, &apiError) {
		switch apiError.(type) {
		case *types.NotFound:
			err = nil
		}
	}

	return false, err
}
