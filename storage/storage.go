package storage

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Storage struct {
	session *session.Session
	client  *s3.S3
}

func NewStorageSession() Storage {
	ss := Storage{}
	sess := session.Must(session.NewSession())
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	ss.session = sess
	ss.client = svc
	return ss
}

func (s Storage) GetBucketObjectList() {
	bucketName := os.Getenv("BUCKET_NAME")
	i := 0
	err := s.client.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &bucketName,
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		i++

		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
}

func (s Storage) GetBucketList() (resp *s3.ListBucketsOutput) {
	resp, err := s.client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	return resp
}
