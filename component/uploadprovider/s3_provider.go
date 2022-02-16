package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"hfs_backend/common"
	"log"
	"net/http"
)

type s3provider struct {
	bucketName string
	region     string
	apikey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketname string, region string, apikey string, secret string, domain string) *s3provider {
	provider := &s3provider{
		bucketName: bucketname,
		region:     region,
		apikey:     apikey,
		secret:     secret,
		domain:     domain,
	}
	s3session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apikey,
			provider.secret,
			""),
	})
	if err != nil {
		log.Fatal(err)
	}
	provider.session = s3session
	return provider
}
func (provider *s3provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileByte := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileByte,
	})
	if err != nil {
		return nil, err
	}
	img := &common.Image{
		Url:       fmt.Sprintf("%s%s", provider.domain, dst),
		CloudName: "s3",
	}
	return img, nil
}
