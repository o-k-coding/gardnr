package objectstorage

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	grdnrconfig "okcoding.com/grdnr/internal/config"
)

type ObjectStorage interface {
	UploadFile(ctx context.Context, key string, contentType string, file io.Reader, expires time.Duration) error // TODO can I stream a file?
	CheckFileExists(ctx context.Context, key string) (bool, error)
	DeleteFile(ctx context.Context, key string) error
}

// TODO create an aws client - and use if config specifies aws or cloudflare for files
// but use the cloudflare client for creating the bucket? idk I guess each post dir should have a bucket?
func NewObjectStorage(ctx context.Context, config grdnrconfig.GrdnrConfig) (ObjectStorage, error) {
	// TODO add config for which type of object storage to use
	// client, err := cloudflare.NewCloudflareClient(config.CloudflareConfig)
	// if err != nil {
	// 	return nil, err
	// }
	// return client, err
	// Currently we use the AWS SDK for Cloudflare R2
	return newAWSS3Client(ctx, config)
}

// TODO everything after this should be a new file
type S3ObjectStorage struct {
	Client     *s3.Client
	bucket     string
	putTimeout time.Duration
	getTimeout time.Duration
}

func newAWSS3Client(ctx context.Context, grdnrConfig grdnrconfig.GrdnrConfig) (*S3ObjectStorage, error) {
	// Create custom resolver for R2 endpoint
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", grdnrConfig.CloudlareAccountID),
		}, nil
	})
	// Load AWS config with custom resolver
	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(grdnrConfig.CloudflareAPIKey, grdnrConfig.CloudflareSecretKey, "")),
		config.WithRegion(grdnrConfig.CloudflareRegion),
	)
	if err != nil {
		return nil, err
	}

	// // Create a new S3 client
	s3Client := s3.NewFromConfig(awsConfig)
	return &S3ObjectStorage{
		Client:     s3Client,
		bucket:     grdnrConfig.CloudflareStorageBucket,
		putTimeout: grdnrConfig.PutObjectTimeout,
		getTimeout: grdnrConfig.GetObjectTimeout,
	}, nil
}

func (s *S3ObjectStorage) UploadFile(ctx context.Context, key string, contentType string, file io.Reader, expires time.Duration) error {
	// Create a PutObjectInput with the specified bucket, key, file content, and content type
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType), // Set the content type to image/jpeg (change as needed)
	}
	if expires > 0 {
		input.Expires = aws.Time(time.Now().Add(expires))
	}

	putCtx, cancel := context.WithTimeoutCause(ctx, s.putTimeout, fmt.Errorf("upload file timeout"))

	defer cancel()

	// Upload the file to Cloudflare R2 Storage
	_, err := s.Client.PutObject(putCtx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3ObjectStorage) CheckFileExists(ctx context.Context, key string) (bool, error) {
	// Create a HeadObjectInput with the specified bucket and key
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}
	headCtx, cancel := context.WithTimeoutCause(ctx, s.getTimeout, fmt.Errorf("check file exists timeout"))
	defer cancel()
	// Check if the file exists
	_, err := s.Client.HeadObject(headCtx, input)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode: 404") {
			return false, nil
		}
		log.Printf("error checking if file exists: %v", err)
		return false, err
	}
	return true, nil
}

func (s *S3ObjectStorage) DeleteFile(ctx context.Context, key string) error {
	// Create a DeleteObjectInput with the specified bucket and key
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}
	deleteCtx, cancel := context.WithTimeoutCause(ctx, s.putTimeout, fmt.Errorf("delete file timeout"))
	defer cancel()
	// Delete the file
	_, err := s.Client.DeleteObject(deleteCtx, input)
	if err != nil {
		log.Printf("error deleting file: %v", err)
		return err
	}
	return nil
}
