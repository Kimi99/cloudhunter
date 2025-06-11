package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

// UserWrapper encapsulates user actions used in the examples.
// It contains an IAM service client that is used to perform user actions.
type UserWrapper struct {
	IamClient *iam.Client
}

func (wrapper UserWrapper) ListUsersWrapper(ctx context.Context, maxUsers int32) ([]types.User, error) {
	result, err := wrapper.IamClient.ListUsers(ctx, &iam.ListUsersInput{
		MaxItems: aws.Int32(maxUsers),
	})

	if err != nil {
		log.Fatal(err)
	}

	return result.Users, err
}