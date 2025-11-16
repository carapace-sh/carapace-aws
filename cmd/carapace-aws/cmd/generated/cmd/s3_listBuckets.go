package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listBucketsCmd = &cobra.Command{
	Use:   "list-buckets",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listBucketsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_listBucketsCmd).Standalone()

		s3_listBucketsCmd.Flags().String("bucket-region", "", "Limits the response to buckets that are located in the specified Amazon Web Services Region.")
		s3_listBucketsCmd.Flags().String("continuation-token", "", "`ContinuationToken` indicates to Amazon S3 that the list is being continued on this bucket with a token.")
		s3_listBucketsCmd.Flags().String("max-buckets", "", "Maximum number of buckets to be returned in response.")
		s3_listBucketsCmd.Flags().String("prefix", "", "Limits the response to bucket names that begin with the specified bucket name prefix.")
	})
	s3Cmd.AddCommand(s3_listBucketsCmd)
}
