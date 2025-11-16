package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteBucketReplicationCmd = &cobra.Command{
	Use:   "delete-bucket-replication",
	Short: "This operation deletes an Amazon S3 on Outposts bucket's replication configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteBucketReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteBucketReplicationCmd).Standalone()

		s3control_deleteBucketReplicationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket to delete the replication configuration for.")
		s3control_deleteBucketReplicationCmd.Flags().String("bucket", "", "Specifies the S3 on Outposts bucket to delete the replication configuration for.")
		s3control_deleteBucketReplicationCmd.MarkFlagRequired("account-id")
		s3control_deleteBucketReplicationCmd.MarkFlagRequired("bucket")
	})
	s3controlCmd.AddCommand(s3control_deleteBucketReplicationCmd)
}
