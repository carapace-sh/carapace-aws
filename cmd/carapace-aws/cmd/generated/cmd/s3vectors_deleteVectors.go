package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_deleteVectorsCmd = &cobra.Command{
	Use:   "delete-vectors",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_deleteVectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3vectors_deleteVectorsCmd).Standalone()

		s3vectors_deleteVectorsCmd.Flags().String("index-arn", "", "The ARN of the vector index that contains a vector you want to delete.")
		s3vectors_deleteVectorsCmd.Flags().String("index-name", "", "The name of the vector index that contains a vector you want to delete.")
		s3vectors_deleteVectorsCmd.Flags().String("keys", "", "The keys of the vectors to delete.")
		s3vectors_deleteVectorsCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket that contains the vector index.")
		s3vectors_deleteVectorsCmd.MarkFlagRequired("keys")
	})
	s3vectorsCmd.AddCommand(s3vectors_deleteVectorsCmd)
}
