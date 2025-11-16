package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_getVectorBucketCmd = &cobra.Command{
	Use:   "get-vector-bucket",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_getVectorBucketCmd).Standalone()

	s3vectors_getVectorBucketCmd.Flags().String("vector-bucket-arn", "", "The ARN of the vector bucket to retrieve information about.")
	s3vectors_getVectorBucketCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket to retrieve information about.")
	s3vectorsCmd.AddCommand(s3vectors_getVectorBucketCmd)
}
