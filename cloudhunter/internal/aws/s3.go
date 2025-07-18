package aws

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kimi99/cloudhunter/internal/shared"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

// S3Wrapper encapsulates the Amazon Simple Storage Service (Amazon S3) actions.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket and object actions.
type S3Wrapper struct {
	S3Client *s3.Client
}

func InitializeS3Wrapper(ctx context.Context, region string, profile string, anonymousMode bool) S3Wrapper {
	if anonymousMode {
		var client = s3.New(s3.Options{
			Credentials: aws.AnonymousCredentials{},
			Region:      region,
		})

		return S3Wrapper{S3Client: client}
	}

	cfg, err := shared.GetAWSConfig(ctx, region, profile)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return S3Wrapper{S3Client: client}
}

func (wrapper S3Wrapper) ListS3BucketContent(ctx context.Context, bucket string, prefix string) ([]*shared.S3Node, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	}

	paginator := s3.NewListObjectsV2Paginator(wrapper.S3Client, input)

	var nodes []*shared.S3Node

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, cp := range output.CommonPrefixes {
			name := strings.TrimPrefix(*cp.Prefix, prefix)
			node := &shared.S3Node{
				Name:     name,
				IsFolder: true,
			}

			childNodes, err := wrapper.ListS3BucketContent(ctx, bucket, *cp.Prefix)
			if err != nil {
				var apiErr smithy.APIError
				if errors.As(err, &apiErr) && apiErr.ErrorCode() == "AccessDenied" {
					log.Printf("[!] Skipping forbidden folder: %s\n", *cp.Prefix)
					nodes = append(nodes, node)
					continue
				}
				return nil, err
			}
			node.Children = childNodes
			nodes = append(nodes, node)
		}

		for _, obj := range output.Contents {
			if *obj.Key == prefix {
				continue
			}

			name := strings.TrimPrefix(*obj.Key, prefix)
			node := &shared.S3Node{
				Name:     name,
				IsFolder: false,
			}
			nodes = append(nodes, node)
		}
	}

	return nodes, nil
}

func (wrapper S3Wrapper) ListBuckets(ctx context.Context) ([]types.Bucket, error) {
	var err error
	var output *s3.ListBucketsOutput
	var buckets []types.Bucket
	bucketPaginator := s3.NewListBucketsPaginator(wrapper.S3Client, &s3.ListBucketsInput{})
	for bucketPaginator.HasMorePages() {
		output, err = bucketPaginator.NextPage(ctx)
		if err != nil {
			var apiErr smithy.APIError
			if errors.As(err, &apiErr) && apiErr.ErrorCode() == "AccessDenied" {
				err = apiErr
			} else {
				return nil, err
			}
			break
		} else {
			buckets = append(buckets, output.Buckets...)
		}
	}
	return buckets, err
}

func (wrapper S3Wrapper) DumpBucketWrapper(ctx context.Context, bucketName string, localFolder string) error {
	paginator := s3.NewListObjectsV2Paginator(wrapper.S3Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			log.Printf("Error listing objects in bucket %s: %v", bucketName, err)
			return err
		}

		for _, obj := range page.Contents {
			key := *obj.Key

			if strings.HasSuffix(key, "/") {
				continue
			}

			localPath := filepath.Join(localFolder, key)

			if err := os.MkdirAll(filepath.Dir(localPath), os.ModePerm); err != nil {
				log.Printf("Couldn't create directory for %s: %v", localPath, err)
				return err
			}

			log.Printf("[+] Downloading: s3://%s/%s", bucketName, key)
			resp, err := wrapper.S3Client.GetObject(ctx, &s3.GetObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(key),
			})
			if err != nil {
				var noKey *types.NoSuchKey
				if errors.As(err, &noKey) {
					log.Printf("Key does not exist: %s", key)
					return err
				}
				log.Printf("Error getting object %s: %v", key, err)
				return err
			}
			defer resp.Body.Close()

			outFile, err := os.Create(localPath)
			if err != nil {
				log.Printf("Couldn't create local file %s: %v", localPath, err)
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, resp.Body)
			if err != nil {
				log.Printf("Failed writing to file %s: %v", localPath, err)
			}
		}
	}

	return nil
}
