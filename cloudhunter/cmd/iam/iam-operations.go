package iam

import (
	"context"
	"fmt"
	"log"

	"github.com/Kimi99/cloudhunter/internal/aws"
	"github.com/Kimi99/cloudhunter/internal/shared"
	"github.com/spf13/cobra"
)

var region string
var profile string
var userName string
var policyName string
var groupName string
var roleName string
var ctx = context.TODO()

var EnumUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Retrieve information about IAM Users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Starting IAM User enumeration...")

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

		users, err := wrapper.ListUsersWrapper(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[+] Found following users:")
		for _, user := range users {
			fmt.Printf("%s\n", *user.UserName)
		}
	},
}

var EnumSpecificUserCmd = &cobra.Command{
	Use:   "get-user",
	Short: "Retrieves information about the specified IAM user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retreiving user information...")

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

		user, err := wrapper.GetUserWrapper(ctx, userName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[+] Retrieved user info:")
		fmt.Printf("ARN: %s\nUsername: %s\nCreated date: %v\n", *user.Arn, *user.UserName, user.CreateDate)
	},
}

var EnumAccessKeysCmd = &cobra.Command{
	Use:   "access-keys",
	Short: "Retrieve information about the IAM access keys associated with the specified IAM user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Starting IAM Access Keys enumeration...")

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

		userPolicies, err := wrapper.ListUserPoliciesWrapper(ctx, userName)
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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

		policy, err := wrapper.GetUserPolicyWrapper(ctx, userName, policyName)
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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

		groups, err := wrapper.ListGroupsForUserWrapper(ctx, userName)
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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

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

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

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

var EnumSpecificRoleCmd = &cobra.Command{
	Use:   "get-role",
	Short: "Retrieve information about the specific IAM role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving role information...")

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

		role, err := wrapper.GetRoleWrapper(ctx, roleName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[+] Retrieved information about role:\n Role ARN: %s\n Role name: %s\n Assume role policy document:\n%s", *role.Role.Arn, *role.Role.RoleName, shared.ParseJsonPolicyDocument(*role.Role.AssumeRolePolicyDocument))
	},
}

var EnumRolePoliciesCmd = &cobra.Command{
	Use:   "role-policies",
	Short: "Retrieve the names of the inline policies that are embedded in the specified IAM role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving role policies...")

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

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
	Short: "Retrieves the specified inline policy document that is embedded with the specified IAM role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retreiving role policy document...")

		wrapper := aws.InitializeIamWrapper(ctx, region, profile)

		policyDocument, err := wrapper.GetRolePolicyDocumentWrapper(ctx, roleName, policyName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[+] Found policy document:\n%s", policyDocument)
	},
}

func init() {
	EnumAccessKeysCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumAccessKeysCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")

	EnumUsersCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumUsersCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")

	EnumSpecificUserCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumSpecificUserCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumSpecificUserCmd.Flags().StringVarP(&userName, "username", "u", "", "Username")

	EnumUserPoliciesCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumUserPoliciesCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumUserPoliciesCmd.Flags().StringVarP(&userName, "username", "u", "", "Username")

	EnumUserPolicyDocumentCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumUserPolicyDocumentCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumUserPolicyDocumentCmd.Flags().StringVarP(&policyName, "policy-name", "n", "pn", "Policy name")
	EnumUserPolicyDocumentCmd.Flags().StringVarP(&userName, "username", "u", "", "Username")

	EnumGroupsCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumGroupsCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")

	EnumGroupsForUserCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumGroupsForUserCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumGroupsForUserCmd.Flags().StringVarP(&userName, "username", "u", "", "Username")

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

	EnumSpecificRoleCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumSpecificRoleCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	EnumSpecificRoleCmd.Flags().StringVarP(&roleName, "role-name", "n", "", "Role name")
}
