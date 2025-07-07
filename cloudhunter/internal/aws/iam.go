package aws

import (
	"context"
	"encoding/json"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

// AwsWrapper encapsulates interaction with AWS services.
// It contains an IAM service client that is used to perform interactions.
// https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_code_examples.html
type AwsWrapper struct {
	IamClient *iam.Client
}

func (wrapper AwsWrapper) ListUsersWrapper(ctx context.Context) ([]types.User, error) {
	result, err := wrapper.IamClient.ListUsers(ctx, &iam.ListUsersInput{})

	if err != nil {
		log.Fatal(err)
	}

	return result.Users, err
}

func (wrapper AwsWrapper) ListAccessKeysWrapper(ctx context.Context) ([]types.AccessKeyMetadata, error) {
	result, err := wrapper.IamClient.ListAccessKeys(ctx, &iam.ListAccessKeysInput{})

	if err != nil {
		log.Fatal(err)
	}

	return result.AccessKeyMetadata, err
}

func (wrapper AwsWrapper) ListUserPoliciesWrapper(ctx context.Context, username string) ([]string, error) {
	result, err := wrapper.IamClient.ListUserPolicies(ctx, &iam.ListUserPoliciesInput{
		UserName: aws.String(username),
	})

	if err != nil {
		log.Fatal(err)
	}

	return result.PolicyNames, err
}

func (wrapper AwsWrapper) GetUserPolicyWrapper(ctx context.Context, username string, policyName string) (string, error) {
	policyDocument, err := wrapper.IamClient.GetUserPolicy(ctx, &iam.GetUserPolicyInput{
		UserName:   aws.String(username),
		PolicyName: aws.String(policyName),
	})

	if err != nil {
		log.Fatal(err)
	}

	policy, err := ParseJsonPolicyDocument(*policyDocument.PolicyDocument)

	return policy, err
}

func (wrapper AwsWrapper) ListGroupsWrapper(ctx context.Context) ([]types.Group, error) {
	groups, err := wrapper.IamClient.ListGroups(ctx, &iam.ListGroupsInput{})

	if err != nil {
		log.Fatal(err)
	}

	return groups.Groups, err
}

func (wrapper AwsWrapper) ListGroupsForUserWrapper(ctx context.Context, username string) ([]types.Group, error) {
	group, err := wrapper.IamClient.ListGroupsForUser(ctx, &iam.ListGroupsForUserInput{
		UserName: &username,
	})

	if err != nil {
		log.Fatal(err)
	}

	return group.Groups, err
}

func (wrapper AwsWrapper) GetGroupWrapper(ctx context.Context, groupName string) (*iam.GetGroupOutput, error) {
	group, err := wrapper.IamClient.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: &groupName,
	})

	if err != nil {
		log.Fatal(err)
	}

	return group, err
}

func (wrapper AwsWrapper) ListGroupPoliciesWrapper(ctx context.Context, groupName string) ([]string, error) {
	policies, err := wrapper.IamClient.ListGroupPolicies(ctx, &iam.ListGroupPoliciesInput{
		GroupName: &groupName,
	})

	if err != nil {
		log.Fatal(err)
	}

	return policies.PolicyNames, err
}

func (wrapper AwsWrapper) GetGroupPolicyDocumentWrapper(ctx context.Context, groupName string, policyName string) (string, error) {
	policyDocument, err := wrapper.IamClient.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{
		GroupName:  &groupName,
		PolicyName: &policyName,
	})

	if err != nil {
		log.Fatal(err)
	}

	policy, err := ParseJsonPolicyDocument(*policyDocument.PolicyDocument)

	return policy, err
}

func (wrapper AwsWrapper) ListRolesWrapper(ctx context.Context) ([]types.Role, error) {
	result, err := wrapper.IamClient.ListRoles(ctx, &iam.ListRolesInput{})

	if err != nil {
		log.Fatal(err)
	}

	return result.Roles, err
}

func ParseJsonPolicyDocument(policyData string) (string, error) {
	decodedPolicy, err := url.QueryUnescape(policyData)
	if err != nil {
		log.Fatal(err)
	}

	var policyObj any
	err = json.Unmarshal([]byte(decodedPolicy), &policyObj)
	if err != nil {
		log.Fatal(err)
	}

	policy, err := json.MarshalIndent(policyObj, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(policy), err
}
