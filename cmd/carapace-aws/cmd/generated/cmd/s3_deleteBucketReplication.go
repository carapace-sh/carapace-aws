package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketReplicationCmd = &cobra.Command{
	Use:   "delete-bucket-replication",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketReplicationCmd).Standalone()

	s3_deleteBucketReplicationCmd.Flags().String("bucket", "", "The bucket name.")
	s3_deleteBucketReplicationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteBucketReplicationCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_deleteBucketReplicationCmd)
}
