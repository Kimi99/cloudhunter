package iam

import (
	"context"
	"fmt"
	"log"

	"github.com/Kimi99/cloudhunter/internal/aws"
	"github.com/Kimi99/cloudhunter/internal/shared"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/spf13/cobra"
)

var region string
var profile string
var username string
var policyName string

var EnumUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Enumerate AWS IAM Users",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		
		fmt.Println("[!] Starting IAM User enumeration...")

		wrapper := initializeAwsWrapper(ctx)

		users, err := wrapper.ListUsersWrapper(ctx)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			fmt.Printf("[+] Found user!\n Username: %s; Created Date: %v\n", *user.UserName, user.CreateDate)
		}
	},
}

var EnumAccessKeysCmd = &cobra.Command{
	Use:	"accesskeys",
	Short: 	"Enumerate IAM Access Keys",
	Run:	func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()

		fmt.Println("[!] Starting IAM Access Keys enumeration...")

		wrapper := initializeAwsWrapper(ctx)

		accessKeys, err := wrapper.ListAccessKeysWrapper(ctx)
		if err != nil {
			log.Fatal(err)
		}

		for _, accessKey := range accessKeys {
			fmt.Printf("[+] Found access key!\n Access Key Id: %s; Username: %s, Status: %s", *accessKey.AccessKeyId, *accessKey.UserName, accessKey.Status)
		}
	},
}

var EnumUserPoliciesCmd = &cobra.Command{
	Use:	"userpolicies",
	Short: 	"Enumerate IAM user policies",
	Run:	func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()

		fmt.Println("[+] Starting IAM user policies enumeration...")

		wrapper := initializeAwsWrapper(ctx)

		userPolicies, err := wrapper.ListUserPolicies(ctx, username)
		if err != nil {
			log.Fatal(err)
		}

		for _, userPolicy := range userPolicies {
			fmt.Printf("[+] Found user policy!\n %s", userPolicy)
		}
	},
}

var EnumPolicyCmd = &cobra.Command{
	Use: 	"get-user-policy",
	Short: 	"Retrieves the specified inline policy document that is embedded in the specified IAM user",
	Run:	func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()

		fmt.Println("[!] Retrieving policy for IAM user...")

		wrapper := initializeAwsWrapper(ctx)

		policy, err := wrapper.GetUserPolicy(ctx, username, policyName)
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf("[+] Found user policy document!\n %s", policy)
	},
}

func init() {
	EnumUsersCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumUsersCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumAccessKeysCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumAccessKeysCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumUserPoliciesCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumUserPoliciesCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumUserPoliciesCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	EnumPolicyCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumPolicyCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumPolicyCmd.Flags().StringVarP(&policyName, "policy-name", "n", "pn", "Policy name")
	EnumPolicyCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
}

func initializeAwsWrapper(ctx context.Context) aws.AwsWrapper {
	cfg, err := shared.GetAWSConfig(ctx, region, profile)
	if err != nil {
		log.Fatal(err)
	}

	client := iam.NewFromConfig(cfg)
	wrapper := aws.AwsWrapper{IamClient: client}

	return wrapper
}