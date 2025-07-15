package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/Kimi99/cloudhunter/internal/aws"
	"github.com/spf13/cobra"
)

var region string
var profile string
var bucketName string
var anonymousMode bool
var ctx = context.TODO()

var ListBucketCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve contents of S3 bucket, if there is any",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving data from bucket...")

		wrapper := aws.InitializeS3Wrapper(ctx, region, profile, anonymousMode)

		objects, err := wrapper.ListS3Bucket(ctx, bucketName, "")
		if err != nil {
			log.Fatal(err)
		}

		printTree(objects, "  ")
	},
}

func printTree(nodes []*aws.S3Node, indent string) {
	for _, node := range nodes {
		if node.IsFolder {
			fmt.Printf("%s %s\n", indent, node.Name)
			printTree(node.Children, indent+"  ")
		} else {
			fmt.Printf("%s %s\n", indent, node.Name)
		}
	}
}

func init() {
	ListBucketCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	ListBucketCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	ListBucketCmd.Flags().StringVarP(&bucketName, "bucket-name", "b", "", "Name of S3 bucket")
	ListBucketCmd.Flags().BoolVarP(&anonymousMode, "anonymous-mode", "a", true, "Use anonymous authentication (true by default)")
}
