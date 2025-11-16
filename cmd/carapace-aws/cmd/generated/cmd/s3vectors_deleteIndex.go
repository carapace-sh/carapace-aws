package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_deleteIndexCmd = &cobra.Command{
	Use:   "delete-index",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_deleteIndexCmd).Standalone()

	s3vectors_deleteIndexCmd.Flags().String("index-arn", "", "The ARN of the vector index to delete.")
	s3vectors_deleteIndexCmd.Flags().String("index-name", "", "The name of the vector index to delete.")
	s3vectors_deleteIndexCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket that contains the vector index.")
	s3vectorsCmd.AddCommand(s3vectors_deleteIndexCmd)
}
