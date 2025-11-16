package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketLocationCmd = &cobra.Command{
	Use:   "get-bucket-location",
	Short: "Using the `GetBucketLocation` operation is no longer a best practice.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketLocationCmd).Standalone()

	s3_getBucketLocationCmd.Flags().String("bucket", "", "The name of the bucket for which to get the location.")
	s3_getBucketLocationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketLocationCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketLocationCmd)
}
