package main

import (
	"os"
	"log"
	"flag"
	"regexp"
	"errors"
	"path"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	region = "ap-southeast-1"
)

var (
	awsAccessKeyId = flag.String("aws-access-key-id", "", "AWS AccessKeyId")
	awsSecretAccessKey = flag.String("aws-secret-key", "", "AWS SecretKey")
	localPath = flag.String("local-path", "", "Local path")
	s3Path = flag.String("s3-path", "", "S3 path s3://<bucket-name>/<path-without-file-name>")
)

func validateCredentials() error {
	if *awsAccessKeyId == "" {
		return errors.New("aws-access-key-id is required")
	}
	if *awsSecretAccessKey == "" {
		return errors.New("aws-secret-key is required")
	}
	return nil
}

func parseS3Path(s3Path string) (string, string, error) {
	re, err := regexp.Compile(`^s3://([^/]+)/(.*)$`)
	if err != nil {
		return "", "", err
	}
	res := re.FindStringSubmatch(s3Path)
	if res == nil {
		return "", "", errors.New("invalid S3 path")
	}
	return res[1], res[2], nil
}

func uploadFile(localPath string, bucket string, remotePath string) error {
	s3Key := path.Join(remotePath, path.Base(localPath))
	svc := s3.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(*awsAccessKeyId, *awsSecretAccessKey, ""),
		Region: region,
	})
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(s3Key),
		Body: file,
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		return err
	}
	log.Println(awsutil.StringValue(resp))
	return nil
}

func main() {
	flag.Parse()
	err := validateCredentials()
	if err != nil {
		log.Fatalf("invalid credentials: %v", err)
	}
	bucket, path, err := parseS3Path(*s3Path)
	if err != nil {
		log.Fatalf("failed to parse s3Path: %v", err)
	}
	log.Printf("Uploading [%s] to [%s] in bucket [%s]", *localPath, path, bucket)
	err = uploadFile(*localPath, bucket, path)
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}
	log.Printf("Done")
}
