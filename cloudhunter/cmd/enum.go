package cmd

import (
	"fmt"

	"github.com/Kimi99/cloudhunter/internal/aws"
	"github.com/spf13/cobra"
)

var enumCmd = &cobra.Command{
	Use:   "enum",
	Short: "Enumerate AWS environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Enumerating IAM users...")
		aws.ListUsers();
	},
}

func init() {
	rootCmd.AddCommand(enumCmd)
}