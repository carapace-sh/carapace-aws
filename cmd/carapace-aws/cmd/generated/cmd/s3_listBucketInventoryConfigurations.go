package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listBucketInventoryConfigurationsCmd = &cobra.Command{
	Use:   "list-bucket-inventory-configurations",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listBucketInventoryConfigurationsCmd).Standalone()

	s3_listBucketInventoryConfigurationsCmd.Flags().String("bucket", "", "The name of the bucket containing the inventory configurations to retrieve.")
	s3_listBucketInventoryConfigurationsCmd.Flags().String("continuation-token", "", "The marker used to continue an inventory configuration listing that has been truncated.")
	s3_listBucketInventoryConfigurationsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_listBucketInventoryConfigurationsCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_listBucketInventoryConfigurationsCmd)
}
