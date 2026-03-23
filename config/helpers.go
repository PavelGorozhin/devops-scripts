package devops

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 is a helper for interacting with S3
type S3 struct{}

// NewS3 returns a new instance of S3
func NewS3() *S3 {
	return &S3{}
}

// UploadFile uploads a file to S3
func (s *S3) UploadFile(ctx context.Context, bucket, file, key string) error {
	if file == "" || bucket == "" {
		return fmt.Errorf("bucket and file must be specified")
	}

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")}, nil)
	if err != nil {
		return err
	}

	s3Client := s3.New(sess)

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(ioutil.ReadFile(file)),
	}

	_, err = s3Client.PutObjectWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

// DownloadFile downloads a file from S3
func (s *S3) DownloadFile(ctx context.Context, bucket, key string) error {
	if bucket == "" || key == "" {
		return fmt.Errorf("bucket and key must be specified")
	}

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")}, nil)
	if err != nil {
		return err
	}

	s3Client := s3.New(sess)

	// Ensure the directory exists
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		if err := os.MkdirAll("output", 0755); err != nil {
			return err
		}
	}

	out, err := os.Create("output/" + key)
	if err != nil {
		return err
	}

	downloader := &s3.Downloader{Session: sess}
	downloader.Download(out, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	return err
}