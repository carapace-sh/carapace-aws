package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_listIndexesCmd = &cobra.Command{
	Use:   "list-indexes",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_listIndexesCmd).Standalone()

	s3vectors_listIndexesCmd.Flags().String("max-results", "", "The maximum number of items to be returned in the response.")
	s3vectors_listIndexesCmd.Flags().String("next-token", "", "The previous pagination token.")
	s3vectors_listIndexesCmd.Flags().String("prefix", "", "Limits the response to vector indexes that begin with the specified prefix.")
	s3vectors_listIndexesCmd.Flags().String("vector-bucket-arn", "", "The ARN of the vector bucket that contains the vector indexes.")
	s3vectors_listIndexesCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket that contains the vector indexes.")
	s3vectorsCmd.AddCommand(s3vectors_listIndexesCmd)
}
