package s3

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
var bucketName string
var anonymousMode bool
var localFolder string
var ctx = context.TODO()

var ListBucketContentCmd = &cobra.Command{
	Use:   "list-content",
	Short: "Retrieve contents of S3 bucket, if there is any",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving data from bucket...")

		wrapper := aws.InitializeS3Wrapper(ctx, region, profile, anonymousMode)

		objects, err := wrapper.ListS3BucketContent(ctx, bucketName, "")
		if err != nil {
			log.Fatal(err)
		}

		if len(objects) != 0 {
			shared.RenderBucketContent(objects, "  ")
		} else {
			fmt.Println("[-] No content is present in the bucket!")
		}
	},
}

var ListBucketsCmd = &cobra.Command{
	Use:   "buckets",
	Short: "Try to retrieve list of S3 buckets present on account. This requires the following policy action: s3:ListAllMyBuckets.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving list of S3 buckets...")

		wrapper := aws.InitializeS3Wrapper(ctx, region, profile, anonymousMode)

		buckets, err := wrapper.ListBuckets(ctx)
		if err != nil {
			fmt.Println("[-] Account does not have sufficient policies to list buckets (required policy action: s3:ListAllMyBuckets).")
			return
		}

		if len(buckets) != 0 {
			fmt.Println("[+] Found following buckets on the account:")
			for _, bucket := range buckets {
				fmt.Printf("%s\n", *bucket.Name)
			}
		} else {
			fmt.Println("[+] No S3 buckets are present on the account!")
		}
	},
}

var DumpBucketCmd = &cobra.Command{
	Use:   "dump-bucket",
	Short: "Try to retrieve the contents of specified S3 bucket.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[!] Retrieving contents of the bucket...")

		wrapper := aws.InitializeS3Wrapper(ctx, region, profile, anonymousMode)

		err := wrapper.DumpBucketWrapper(ctx, bucketName, localFolder)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("[+] Dumped contents of S3 bucket to local folder: %s", localFolder)
		}
	},
}

func init() {
	ListBucketContentCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	ListBucketContentCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	ListBucketContentCmd.Flags().StringVarP(&bucketName, "bucket-name", "b", "", "Name of S3 bucket")
	ListBucketContentCmd.Flags().BoolVarP(&anonymousMode, "anonymous-mode", "a", false, "Use anonymous authentication")

	ListBucketsCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	ListBucketsCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")

	DumpBucketCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	DumpBucketCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile")
	DumpBucketCmd.Flags().StringVarP(&bucketName, "bucket-name", "b", "", "Name of S3 bucket")
	DumpBucketCmd.Flags().BoolVarP(&anonymousMode, "anonymous-mode", "a", false, "Use anonymous authentication")
	DumpBucketCmd.Flags().StringVarP(&localFolder, "folder", "f", "bucket", "Local folder used to store the bucket content")
}
