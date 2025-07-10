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
var groupName string
var roleName string
var ctx = context.TODO()

var EnumUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Retrieve information about IAM Users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Starting IAM User enumeration...")

		wrapper := initializeAwsWrapper(ctx)

		users, err := wrapper.ListUsersWrapper(ctx)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			fmt.Printf("[+] Found user!\n Username: %s\n Created Date: %v\n", *user.UserName, user.CreateDate)
		}
	},
}

var EnumAccessKeysCmd = &cobra.Command{
	Use:   "access-keys",
	Short: "Retrieve information about the IAM access keys associated with the specified IAM user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Starting IAM Access Keys enumeration...")

		wrapper := initializeAwsWrapper(ctx)

		accessKeys, err := wrapper.ListAccessKeysWrapper(ctx)
		if err != nil {
			log.Fatal(err)
		}

		for _, accessKey := range accessKeys {
			fmt.Printf("[+] Found access key!\n Access Key Id: %s\n Username: %s\n Status: %s", *accessKey.AccessKeyId, *accessKey.UserName, accessKey.Status)
		}
	},
}

var EnumUserPoliciesCmd = &cobra.Command{
	Use:   "user-policies",
	Short: "Retrieve names of the inline policies embedded in the specified IAM user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[+] Starting IAM user policies enumeration...")

		wrapper := initializeAwsWrapper(ctx)

		userPolicies, err := wrapper.ListUserPoliciesWrapper(ctx, username)
		if err != nil {
			log.Fatal(err)
		}

		for _, userPolicy := range userPolicies {
			fmt.Printf("[+] Found user policy!\n %s", userPolicy)
		}
	},
}

var EnumUserPolicyDocumentCmd = &cobra.Command{
	Use:   "get-user-policy-document",
	Short: "Retrieves the specified inline policy document that is embedded in the specified IAM user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving policy document for IAM user...")

		wrapper := initializeAwsWrapper(ctx)

		policy, err := wrapper.GetUserPolicyWrapper(ctx, username, policyName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[+] Found user policy document!\n %s", policy)
	},
}

var EnumGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Retrieve information about IAM groups",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving groups from IAM...")

		wrapper := initializeAwsWrapper(ctx)

		groups, err := wrapper.ListGroupsWrapper(ctx)
		if err != nil {
			log.Fatal(err)
		}

		for _, group := range groups {
			fmt.Printf("[+] Found group!\n Group ARN: %s\n Group name: %s\n", *group.Arn, *group.GroupName)
		}
	},
}

var EnumGroupsForUserCmd = &cobra.Command{
	Use:   "user-groups",
	Short: "Retrieve information about the IAM groups that the specified IAM user belongs to",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving user groups from IAM...")

		wrapper := initializeAwsWrapper(ctx)

		groups, err := wrapper.ListGroupsForUserWrapper(ctx, username)
		if err != nil {
			log.Fatal(err)
		}

		for _, group := range groups {
			fmt.Printf("[+] Found group!\n Group ARN: %s\n Group name: %s\n", *group.Arn, *group.GroupName)
		}
	},
}

var EnumSpecificGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Retrieve information about the specific IAM group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving information about the specified IAM group...")

		wrapper := initializeAwsWrapper(ctx)

		group, err := wrapper.GetGroupWrapper(ctx, groupName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[+] Retrieved information about group:\n Group ARN: %s\n Group name: %s\n", *group.Group.Arn, *group.Group.GroupName)
		fmt.Println("\n[+] Listing out group users...")
		for _, user := range group.Users {
			fmt.Printf("Username: %s\nCreated Date: %v\n", *user.UserName, user.CreateDate)
			fmt.Println()
		}
	},
}

var EnumGroupPoliciesCmd = &cobra.Command{
	Use:   "group-policies",
	Short: "Retrieve the names of the inline policies that are embedded in the specified IAM group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving group policies...")

		wrapper := initializeAwsWrapper(ctx)

		policies, err := wrapper.ListGroupPoliciesWrapper(ctx, groupName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[+] Found following group policies:")
		for _, policy := range policies {
			fmt.Printf("%s\n", policy)
		}
	},
}

var EnumGroupPolicyDocumentCmd = &cobra.Command{
	Use:   "get-group-policy-document",
	Short: "Retrieves the specified inline policy document that is embedded in the specified IAM group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving group policy document...")

		wrapper := initializeAwsWrapper(ctx)

		policyDocument, err := wrapper.GetGroupPolicyDocumentWrapper(ctx, groupName, policyName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[+] Found policy document:\n%s", policyDocument)
	},
}

var EnumRolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "Retrieves the list of IAM roles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving roles...")

		wrapper := initializeAwsWrapper(ctx)

		roles, err := wrapper.ListRolesWrapper(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[+] Found following roles:")
		for _, role := range roles {
			fmt.Printf("%s\n", *role.RoleName)
		}
	},
}

var EnumRolePoliciesCmd = &cobra.Command{
	Use:   "role-policies",
	Short: "Lists the names of the inline policies that are embedded in the specified IAM role.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving role policies...")

		wrapper := initializeAwsWrapper(ctx)

		policies, err := wrapper.ListRolePoliciesWrapper(ctx, roleName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[+] Found following role policies:")
		for _, policy := range policies {
			fmt.Printf("%s", policy)
		}
	},
}

var EnumRolePolicyDocumentCmd = &cobra.Command{
	Use:   "get-role-policy-document",
	Short: "Retrieves the specified inline policy document that is embedded with the specified IAM role.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retreiving role policy document...")

		wrapper := initializeAwsWrapper(ctx)

		policyDocument, err := wrapper.GetRolePolicyDocumentWrapper(ctx, roleName, policyName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[+] Found policy document:\n%s", policyDocument)
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

	EnumUserPolicyDocumentCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumUserPolicyDocumentCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumUserPolicyDocumentCmd.Flags().StringVarP(&policyName, "policy-name", "n", "pn", "Policy name")
	EnumUserPolicyDocumentCmd.Flags().StringVarP(&username, "username", "u", "", "Username")

	EnumGroupsCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumGroupsCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")

	EnumGroupsForUserCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumGroupsForUserCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumGroupsForUserCmd.Flags().StringVarP(&username, "username", "u", "", "Username")

	EnumSpecificGroupCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumSpecificGroupCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumSpecificGroupCmd.Flags().StringVarP(&groupName, "groupname", "g", "", "Group name")

	EnumGroupPoliciesCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumGroupPoliciesCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumGroupPoliciesCmd.Flags().StringVarP(&groupName, "groupname", "g", "", "Group name")

	EnumGroupPolicyDocumentCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumGroupPolicyDocumentCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumGroupPolicyDocumentCmd.Flags().StringVarP(&policyName, "policy-name", "n", "", "Policy name")
	EnumGroupPolicyDocumentCmd.Flags().StringVarP(&groupName, "groupname", "g", "", "Group name")

	EnumRolesCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumRolesCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")

	EnumRolePoliciesCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumRolePoliciesCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumRolePoliciesCmd.Flags().StringVarP(&roleName, "role-name", "n", "", "Role name")

	EnumRolePolicyDocumentCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumRolePolicyDocumentCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumRolePolicyDocumentCmd.Flags().StringVarP(&policyName, "policy-name", "n", "", "Policy name")
	EnumRolePolicyDocumentCmd.Flags().StringVarP(&roleName, "rolename", "l", "", "Role name")
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
