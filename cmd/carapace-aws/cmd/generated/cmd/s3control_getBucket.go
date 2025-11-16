package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getBucketCmd = &cobra.Command{
	Use:   "get-bucket",
	Short: "Gets an Amazon S3 on Outposts bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getBucketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getBucketCmd).Standalone()

		s3control_getBucketCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
		s3control_getBucketCmd.Flags().String("bucket", "", "Specifies the bucket.")
		s3control_getBucketCmd.MarkFlagRequired("account-id")
		s3control_getBucketCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_getBucketCmd)
}
