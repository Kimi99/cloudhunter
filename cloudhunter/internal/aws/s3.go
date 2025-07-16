package aws

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/Kimi99/cloudhunter/internal/shared"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

// Function that retrieves objects from S3 bucket. First folders and then recursively files from each folder it has access to
func (wrapper S3Wrapper) ListS3Bucket(ctx context.Context, bucket string, prefix string) ([]*shared.S3Node, error) {
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

			childNodes, err := wrapper.ListS3Bucket(ctx, bucket, *cp.Prefix)
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
