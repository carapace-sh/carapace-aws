package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getBucketReplicationCmd = &cobra.Command{
	Use:   "get-bucket-replication",
	Short: "This operation gets an Amazon S3 on Outposts bucket's replication configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getBucketReplicationCmd).Standalone()

	s3control_getBucketReplicationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
	s3control_getBucketReplicationCmd.Flags().String("bucket", "", "Specifies the bucket to get the replication information for.")
	s3control_getBucketReplicationCmd.MarkFlagRequired("account-id")
	s3control_getBucketReplicationCmd.MarkFlagRequired("bucket")
	s3controlCmd.AddCommand(s3control_getBucketReplicationCmd)
}
