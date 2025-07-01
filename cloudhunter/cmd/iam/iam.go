package iam

import "github.com/spf13/cobra"

var IamCmd = &cobra.Command {
	Use: "iam",
	Short: "Interact with AWS IAM service",
}