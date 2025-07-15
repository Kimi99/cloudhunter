package aws

import (
	"context"
	"log"

	"github.com/Kimi99/cloudhunter/internal/shared"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

// AwsWrapper encapsulates interaction with AWS services.
// It contains an IAM service client that is used to perform interactions.
// https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_code_examples.html
type IamWrapper struct {
	IamClient *iam.Client
}

func (wrapper IamWrapper) ListAccessKeysWrapper(ctx context.Context) ([]types.AccessKeyMetadata, error) {
	accessKeys, err := wrapper.IamClient.ListAccessKeys(ctx, &iam.ListAccessKeysInput{})

	if err != nil {
		log.Fatal(err)
	}

	return accessKeys.AccessKeyMetadata, err
}

func (wrapper IamWrapper) ListUsersWrapper(ctx context.Context) ([]types.User, error) {
	users, err := wrapper.IamClient.ListUsers(ctx, &iam.ListUsersInput{})

	if err != nil {
		log.Fatal(err)
	}

	return users.Users, err
}

func (wrapper IamWrapper) GetUserWrapper(ctx context.Context, userName string) (types.User, error) {
	user, err := wrapper.IamClient.GetUser(ctx, &iam.GetUserInput{
		UserName: &userName,
	})

	if err != nil {
		log.Fatal(err)
	}

	return *user.User, err
}

func (wrapper IamWrapper) ListUserPoliciesWrapper(ctx context.Context, username string) ([]string, error) {
	policies, err := wrapper.IamClient.ListUserPolicies(ctx, &iam.ListUserPoliciesInput{
		UserName: aws.String(username),
	})

	if err != nil {
		log.Fatal(err)
	}

	return policies.PolicyNames, err
}

func (wrapper IamWrapper) GetUserPolicyWrapper(ctx context.Context, username string, policyName string) (string, error) {
	policyDocument, err := wrapper.IamClient.GetUserPolicy(ctx, &iam.GetUserPolicyInput{
		UserName:   aws.String(username),
		PolicyName: aws.String(policyName),
	})

	if err != nil {
		log.Fatal(err)
	}

	policy := shared.ParseJsonPolicyDocument(*policyDocument.PolicyDocument)

	return policy, err
}

func (wrapper IamWrapper) ListGroupsWrapper(ctx context.Context) ([]types.Group, error) {
	groups, err := wrapper.IamClient.ListGroups(ctx, &iam.ListGroupsInput{})

	if err != nil {
		log.Fatal(err)
	}

	return groups.Groups, err
}

func (wrapper IamWrapper) ListGroupsForUserWrapper(ctx context.Context, username string) ([]types.Group, error) {
	group, err := wrapper.IamClient.ListGroupsForUser(ctx, &iam.ListGroupsForUserInput{
		UserName: &username,
	})

	if err != nil {
		log.Fatal(err)
	}

	return group.Groups, err
}

func (wrapper IamWrapper) GetGroupWrapper(ctx context.Context, groupName string) (*iam.GetGroupOutput, error) {
	group, err := wrapper.IamClient.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: &groupName,
	})

	if err != nil {
		log.Fatal(err)
	}

	return group, err
}

func (wrapper IamWrapper) ListGroupPoliciesWrapper(ctx context.Context, groupName string) ([]string, error) {
	policies, err := wrapper.IamClient.ListGroupPolicies(ctx, &iam.ListGroupPoliciesInput{
		GroupName: &groupName,
	})

	if err != nil {
		log.Fatal(err)
	}

	return policies.PolicyNames, err
}

func (wrapper IamWrapper) GetGroupPolicyDocumentWrapper(ctx context.Context, groupName string, policyName string) (string, error) {
	policyDocument, err := wrapper.IamClient.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{
		GroupName:  &groupName,
		PolicyName: &policyName,
	})

	if err != nil {
		log.Fatal(err)
	}

	policy := shared.ParseJsonPolicyDocument(*policyDocument.PolicyDocument)

	return policy, err
}

func (wrapper IamWrapper) ListRolesWrapper(ctx context.Context) ([]types.Role, error) {
	result, err := wrapper.IamClient.ListRoles(ctx, &iam.ListRolesInput{})

	if err != nil {
		log.Fatal(err)
	}

	return result.Roles, err
}

func (wrapper IamWrapper) GetRoleWrapper(ctx context.Context, roleName string) (*iam.GetRoleOutput, error) {
	role, err := wrapper.IamClient.GetRole(ctx, &iam.GetRoleInput{
		RoleName: &roleName,
	})

	if err != nil {
		log.Fatal(err)
	}

	return role, err
}

func (wrapper IamWrapper) ListRolePoliciesWrapper(ctx context.Context, roleName string) ([]string, error) {
	result, err := wrapper.IamClient.ListRolePolicies(ctx, &iam.ListRolePoliciesInput{
		RoleName: &roleName,
	})

	if err != nil {
		log.Fatal(err)
	}

	return result.PolicyNames, err
}

func (wrapper IamWrapper) GetRolePolicyDocumentWrapper(ctx context.Context, roleName string, policyName string) (string, error) {
	policyDocument, err := wrapper.IamClient.GetRolePolicy(ctx, &iam.GetRolePolicyInput{
		RoleName:   &roleName,
		PolicyName: &policyName,
	})

	if err != nil {
		log.Fatal(err)
	}

	policy := shared.ParseJsonPolicyDocument(*policyDocument.PolicyDocument)

	return policy, err
}

func InitializeIamWrapper(ctx context.Context, region string, profile string) IamWrapper {
	cfg, err := shared.GetAWSConfig(ctx, region, profile)
	if err != nil {
		log.Fatal(err)
	}

	client := iam.NewFromConfig(cfg)
	wrapper := IamWrapper{IamClient: client}

	return wrapper
}
