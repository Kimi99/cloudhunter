package cmd

import (
	"context"
	"fmt"
	"log"
	"slices"

	"github.com/Kimi99/cloudhunter/internal/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/spf13/cobra"
)

var enumCmd = &cobra.Command{
	Use:   "iam",
	Short: "Enumerate AWS IAM",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Enumerating IAM users...")
		
		ctx := context.TODO()

		fmt.Println("[!] Enter AWS account region: ")

		var region string
		_, err := fmt.Scanln(&region)

		regions := []string { 
			"ap-northeast-1",
			"ap-northeast-2",
			"ap-northeast-3",
			"ap-south-1",
			"ap-southeast-1",
			"ap-southeast-2",
			"ca-central-1",
			"eu-central-1",
			"eu-north-1",
			"eu-west-1",
			"eu-west-2",
			"eu-west-3",
			"sa-east-1",
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2" }

		if err != nil  {
			log.Fatal(err)
		} else if !slices.Contains(regions, region) {
			fmt.Println("[-] Region not valid.")
			return
		}

		cfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion(region),
		)
		if err != nil {
			log.Fatal(err)
		}

		client := iam.NewFromConfig(cfg)
		wrapper := aws.UserWrapper{IamClient: client}

		users, err := wrapper.ListUsersWrapper(ctx, 10)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			fmt.Printf("[+] Found user! Username: %s; Created Date: %v\n", *user.UserName, user.CreateDate)
		}
	},
}

func init() {
	rootCmd.AddCommand(enumCmd)
}