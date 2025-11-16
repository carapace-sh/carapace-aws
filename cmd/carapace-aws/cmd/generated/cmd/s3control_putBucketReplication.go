package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putBucketReplicationCmd = &cobra.Command{
	Use:   "put-bucket-replication",
	Short: "This action creates an Amazon S3 on Outposts bucket's replication configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putBucketReplicationCmd).Standalone()

	s3control_putBucketReplicationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
	s3control_putBucketReplicationCmd.Flags().String("bucket", "", "Specifies the S3 on Outposts bucket to set the configuration for.")
	s3control_putBucketReplicationCmd.Flags().String("replication-configuration", "", "")
	s3control_putBucketReplicationCmd.MarkFlagRequired("account-id")
	s3control_putBucketReplicationCmd.MarkFlagRequired("bucket")
	s3control_putBucketReplicationCmd.MarkFlagRequired("replication-configuration")
	s3controlCmd.AddCommand(s3control_putBucketReplicationCmd)
}
