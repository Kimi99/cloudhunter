package cmd

import (
	"github.com/Kimi99/cloudhunter/cmd/iam"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cloudhunter",
	Short: "CloudHunter - AWS post-compromise enumeration tool",
	Long: `CloudHunter is a CLI tool for mapping AWS environments using 
stolen or assumed credentials in post-compromise or red team scenarios.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	iam.IamCmd.AddCommand(iam.EnumUsersCmd)
	rootCmd.AddCommand(iam.IamCmd)
}