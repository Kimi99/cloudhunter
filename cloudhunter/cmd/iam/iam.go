package iam

import "github.com/spf13/cobra"

var IamCmd = &cobra.Command {
	Use: "iam",
	Short: "Interact with AWS IAM service",
}

func init() {
	IamCmd.AddCommand(EnumUsersCmd)
	
	IamCmd.AddCommand(EnumAccessKeysCmd)
	
	IamCmd.AddCommand(EnumUserPoliciesCmd)
	IamCmd.AddCommand(EnumUserPolicyDocumentCmd)

	IamCmd.AddCommand(EnumGroupsCmd)
	IamCmd.AddCommand(EnumGroupsForUserCmd)
}