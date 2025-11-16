package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getPublicAccessBlockCmd = &cobra.Command{
	Use:   "get-public-access-block",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getPublicAccessBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getPublicAccessBlockCmd).Standalone()

		s3_getPublicAccessBlockCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose `PublicAccessBlock` configuration you want to retrieve.")
		s3_getPublicAccessBlockCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getPublicAccessBlockCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getPublicAccessBlockCmd)
}
