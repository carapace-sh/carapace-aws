package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_listVectorBucketsCmd = &cobra.Command{
	Use:   "list-vector-buckets",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_listVectorBucketsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3vectors_listVectorBucketsCmd).Standalone()

		s3vectors_listVectorBucketsCmd.Flags().String("max-results", "", "The maximum number of vector buckets to be returned in the response.")
		s3vectors_listVectorBucketsCmd.Flags().String("next-token", "", "The previous pagination token.")
		s3vectors_listVectorBucketsCmd.Flags().String("prefix", "", "Limits the response to vector buckets that begin with the specified prefix.")
	})
	s3vectorsCmd.AddCommand(s3vectors_listVectorBucketsCmd)
}
