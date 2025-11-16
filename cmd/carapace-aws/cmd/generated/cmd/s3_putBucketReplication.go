package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketReplicationCmd = &cobra.Command{
	Use:   "put-bucket-replication",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketReplicationCmd).Standalone()

	s3_putBucketReplicationCmd.Flags().String("bucket", "", "The name of the bucket")
	s3_putBucketReplicationCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
	s3_putBucketReplicationCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the data.")
	s3_putBucketReplicationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putBucketReplicationCmd.Flags().String("replication-configuration", "", "")
	s3_putBucketReplicationCmd.Flags().String("token", "", "A token to allow Object Lock to be enabled for an existing bucket.")
	s3_putBucketReplicationCmd.MarkFlagRequired("bucket")
	s3_putBucketReplicationCmd.MarkFlagRequired("replication-configuration")
	s3Cmd.AddCommand(s3_putBucketReplicationCmd)
}
