package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_createVectorBucketCmd = &cobra.Command{
	Use:   "create-vector-bucket",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_createVectorBucketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3vectors_createVectorBucketCmd).Standalone()

		s3vectors_createVectorBucketCmd.Flags().String("encryption-configuration", "", "The encryption configuration for the vector bucket.")
		s3vectors_createVectorBucketCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket to create.")
		s3vectors_createVectorBucketCmd.MarkFlagRequired("vector-bucket-name")
	})
	s3vectorsCmd.AddCommand(s3vectors_createVectorBucketCmd)
}
