package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

func MakeClient() (*minio.Client, error) {
	endPoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	const useSSL = false

	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return minioClient, nil
}

func BucketExists(bucketName string) (bool, error) {
	client, err := MakeClient()
	if err != nil {
		return false, err
	}
	ctx := context.Background()

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func makeBucket(bucketName string) error {
	client, err := MakeClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}
	return nil
}

func UploadFile(fileName string, filePath string) {
	client, err := MakeClient()
	if err != nil {
		panic(err)
		return
	}
	ctx := context.Background()
	bucketName := os.Getenv("MINIO_BUCKET")

	exists, err := BucketExists(bucketName)
	if err != nil {
		panic(err)
		return
	} else if !exists {
		err = makeBucket(bucketName)
		if err != nil {
			panic(err)
			return
		}
	}

	uploadInfo, err := client.FPutObject(ctx, bucketName, fileName, filePath, minio.PutObjectOptions{})
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("Uploaded %v at %v", fileName, uploadInfo.Location)
}
