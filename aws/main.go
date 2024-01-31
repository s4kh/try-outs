package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const bucket string = "test-sdk-list"
const srcBucket string = "test-sdk-src"

func emptyBucket(s3client *s3.Client, ctx context.Context) error {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}
	p := s3.NewListObjectsV2Paginator(s3client, params)

	for p.HasMorePages() {

		page, err := p.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("error getting the page during empty bucket: %v", err)
		}

		var objKeys []string
		for _, object := range page.Contents {
			objKeys = append(objKeys, *object.Key)
		}
		if len(objKeys) == 0 {
			return nil
		}

		var objectIds []types.ObjectIdentifier
		for _, key := range objKeys {
			objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
		}
		_, err = s3client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(bucket),
			Delete: &types.Delete{Objects: objectIds},
		})

		if err != nil {
			return fmt.Errorf("error deleting objects from bucket %q: %v", bucket, err)
		}
	}

	return nil
}

func uploadToAWS(client *s3.Client) error {
	p := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})

	fmt.Println("list resutl:")

	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}

		for _, obj := range page.Contents {
			fmt.Println("Object: ", *obj.Key)
		}
	}

	return nil
}

func initAWS() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("dev"))
	if err != nil {
		return nil, fmt.Errorf("error during aws config load: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return client, nil
}

func populateBucket(s3client *s3.Client, targetBucket, sourceBucket string, ctx context.Context) error {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(sourceBucket),
	}
	p := s3.NewListObjectsV2Paginator(s3client, params)

	for p.HasMorePages() {

		page, err := p.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("error getting the page during populate bucket: %v", err)
		}

		for _, object := range page.Contents {
			fmt.Printf("%v/%v", sourceBucket, *object.Key)
			_, err := s3client.CopyObject(ctx, &s3.CopyObjectInput{
				Bucket:     aws.String(targetBucket),
				CopySource: aws.String(fmt.Sprintf("%v/%v", sourceBucket, *object.Key)),
				Key:        object.Key,
			})

			if err != nil {
				return fmt.Errorf("failed to copy object %q: %v", *object.Key, err)
			}
		}
	}

	return nil
}

func main() {
	client, err := initAWS()

	if err != nil {
		log.Fatal(err)
	}

	err = emptyBucket(client, context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = populateBucket(client, bucket, srcBucket, context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
