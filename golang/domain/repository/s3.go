package repository

type S3Repository interface {
	CreateBucket(bucketName string) error
}
