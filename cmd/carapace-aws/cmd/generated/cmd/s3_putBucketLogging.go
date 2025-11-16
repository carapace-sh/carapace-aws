package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketLoggingCmd = &cobra.Command{
	Use:   "put-bucket-logging",
	Short: "End of support notice: As of October 1, 2025, Amazon S3 has discontinued support for Email Grantee Access Control Lists (ACLs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketLoggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketLoggingCmd).Standalone()

		s3_putBucketLoggingCmd.Flags().String("bucket", "", "The name of the bucket for which to set the logging parameters.")
		s3_putBucketLoggingCmd.Flags().String("bucket-logging-status", "", "Container for logging status information.")
		s3_putBucketLoggingCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketLoggingCmd.Flags().String("content-md5", "", "The MD5 hash of the `PutBucketLogging` request body.")
		s3_putBucketLoggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketLoggingCmd.MarkFlagRequired("bucket")
		s3_putBucketLoggingCmd.MarkFlagRequired("bucket-logging-status")
	})
	s3Cmd.AddCommand(s3_putBucketLoggingCmd)
}
