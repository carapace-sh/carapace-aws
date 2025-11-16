package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketReplicationCmd = &cobra.Command{
	Use:   "get-bucket-replication",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketReplicationCmd).Standalone()

	s3_getBucketReplicationCmd.Flags().String("bucket", "", "The bucket name for which to get the replication information.")
	s3_getBucketReplicationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketReplicationCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketReplicationCmd)
}
