package s3

import "github.com/spf13/cobra"

var S3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Interact with AWS S3 service",
}

func init() {
	S3Cmd.AddCommand(ListBucketContentCmd)
	S3Cmd.AddCommand(ListBucketsCmd)
	S3Cmd.AddCommand(DumpBucketCmd)
}
