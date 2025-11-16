package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listDirectoryBucketsCmd = &cobra.Command{
	Use:   "list-directory-buckets",
	Short: "Returns a list of all Amazon S3 directory buckets owned by the authenticated sender of the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listDirectoryBucketsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_listDirectoryBucketsCmd).Standalone()

		s3_listDirectoryBucketsCmd.Flags().String("continuation-token", "", "`ContinuationToken` indicates to Amazon S3 that the list is being continued on buckets in this account with a token.")
		s3_listDirectoryBucketsCmd.Flags().String("max-directory-buckets", "", "Maximum number of buckets to be returned in response.")
	})
	s3Cmd.AddCommand(s3_listDirectoryBucketsCmd)
}
