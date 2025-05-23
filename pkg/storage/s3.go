package storage

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/Yobubble/yona-bot/config"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// Ref: https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/sdk-utilities-s3.html
type s3Storage struct {
	cli   *s3.Client
	s3Cfg *config.S3Config
}

func (s *s3Storage) Read(filePath string) ([]byte, error) {
	trimmed := strings.TrimPrefix(filePath, "./")

	downloader := manager.NewDownloader(s.cli)
	buf := manager.NewWriteAtBuffer([]byte{})

	_, err := downloader.Download(context.TODO(), buf, &s3.GetObjectInput{
		Bucket: aws.String(s.s3Cfg.S3Bucket),
		Key:    aws.String(trimmed),
	})

	if err != nil {
		return nil, fmt.Errorf("storage: failed to download file '%s' from bucket '%s': %w", filePath, s.s3Cfg.S3Bucket, err)
	}

	return buf.Bytes(), nil
}

func (s *s3Storage) Write(data []byte, filePath string) error {
	trimmed := strings.TrimPrefix(filePath, "./")

	uploader := manager.NewUploader(s.cli)
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.s3Cfg.S3Bucket),
		Key:    aws.String(trimmed),
		Body:   bytes.NewBuffer(data),
	})
	if err != nil {
		return fmt.Errorf("storage: failed to upload file '%s' to bucket '%s': %w", filePath, s.s3Cfg.S3Bucket, err)
	}

	return nil
}

func newS3(s3Cfg *config.S3Config) Storage {
	cfg, _ := awsConfig.LoadDefaultConfig(context.TODO())

	cli := s3.NewFromConfig(cfg)

	return &s3Storage{
		cli:   cli,
		s3Cfg: s3Cfg,
	}
}
