package storage

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

type Storage struct {
	session *session.Session
	client  *s3.S3
}

func NewStorageSession() Storage {
	ss := Storage{}
	sess := session.Must(session.NewSession())
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(viper.GetString("aws_region")),
	})
	ss.session = sess
	ss.client = svc
	return ss
}

func (s Storage) GetBucketObjectList(bucketName string) (resp *s3.ListObjectsV2Output) {
	resp, err := s.client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		panic(err)
	}

	return resp
}

func (s Storage) GetBucketList() (resp *s3.ListBucketsOutput) {
	resp, err := s.client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	return resp
}

type UploadRequest struct {
	FileName   string `json:"file_name"`
	BucketName string `json:"bucket_name"`
	ObjectKey  string `json:"object_key"`
}

func (s Storage) UploadFile(r UploadRequest) (location string, err error) {
	file, err := os.Open(r.FileName)
	if err != nil {
		return "", nil
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return "", nil
	}

	reader := &CustomReader{
		fp:      file,
		size:    fileInfo.Size(),
		signMap: map[int64]struct{}{},
	}

	uploader := s3manager.NewUploader(s.session, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024
		u.LeavePartsOnError = true
	})

	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(r.BucketName),
		Key:    aws.String(r.ObjectKey),
		Body:   reader,
	})
	if err != nil {
		return "", err
	}

	return output.Location, nil
}
