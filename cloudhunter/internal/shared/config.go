package shared

import (
	"context"
	"fmt"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var ValidRegions = []string{
	"ap-northeast-1", "ap-northeast-2", "ap-northeast-3",
	"ap-south-1", "ap-southeast-1", "ap-southeast-2",
	"ca-central-1", "eu-central-1", "eu-north-1",
	"eu-west-1", "eu-west-2", "eu-west-3",
	"sa-east-1", "us-east-1", "us-east-2",
	"us-west-1", "us-west-2",
}

func GetAWSConfig(ctx context.Context, region, profile string) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}

	if profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(profile))
	}

	if region != "" {
		if !slices.Contains(ValidRegions, region) {
			return aws.Config{}, fmt.Errorf("[-] Invalid AWS region: %s", region)
		}
		opts = append(opts, config.WithRegion(region))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return cfg, fmt.Errorf("[-] Failed to load AWS config: %w", err)
	}

	return cfg, nil
}