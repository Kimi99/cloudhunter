package iam

import "github.com/spf13/cobra"

var IamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Interact with AWS IAM service",
}

func init() {
	IamCmd.AddCommand(EnumUsersCmd)
	IamCmd.AddCommand(EnumSpecificUserCmd)

	IamCmd.AddCommand(EnumAccessKeysCmd)

	IamCmd.AddCommand(EnumUserPoliciesCmd)
	IamCmd.AddCommand(EnumUserPolicyDocumentCmd)

	IamCmd.AddCommand(EnumGroupsCmd)
	IamCmd.AddCommand(EnumGroupsForUserCmd)
	IamCmd.AddCommand(EnumSpecificGroupCmd)
	IamCmd.AddCommand(EnumGroupPoliciesCmd)
	IamCmd.AddCommand(EnumGroupPolicyDocumentCmd)

	IamCmd.AddCommand(EnumRolesCmd)
	IamCmd.AddCommand(EnumRolePoliciesCmd)
	IamCmd.AddCommand(EnumRolePolicyDocumentCmd)
	IamCmd.AddCommand(EnumSpecificRoleCmd)
}
