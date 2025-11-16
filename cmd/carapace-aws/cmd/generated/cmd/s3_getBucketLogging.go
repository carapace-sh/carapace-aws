package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketLoggingCmd = &cobra.Command{
	Use:   "get-bucket-logging",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketLoggingCmd).Standalone()

	s3_getBucketLoggingCmd.Flags().String("bucket", "", "The bucket name for which to get the logging information.")
	s3_getBucketLoggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketLoggingCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketLoggingCmd)
}
