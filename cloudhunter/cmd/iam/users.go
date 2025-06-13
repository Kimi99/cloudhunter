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

var EnumUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Enumerate AWS IAM Users",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		
		fmt.Println("[!] Starting IAM User enumeration...")

		cfg, err := shared.GetAWSConfig(ctx, region, profile)
		if err != nil {
			log.Fatal(err)
		}

		client := iam.NewFromConfig(cfg)
		wrapper := aws.UserWrapper{IamClient: client}

		users, err := wrapper.ListUsersWrapper(ctx, 100)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			fmt.Printf("[+] Found user! Username: %s; Created Date: %v\n", *user.UserName, user.CreateDate)
		}
	},
}

func init() {
	EnumUsersCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	EnumUsersCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
}