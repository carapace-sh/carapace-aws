package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketMetadataTableConfigurationCmd = &cobra.Command{
	Use:   "delete-bucket-metadata-table-configuration",
	Short: "We recommend that you delete your S3 Metadata configurations by using the V2 [DeleteBucketMetadataTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataTableConfiguration.html) API operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketMetadataTableConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteBucketMetadataTableConfigurationCmd).Standalone()

		s3_deleteBucketMetadataTableConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that you want to remove the metadata table configuration from.")
		s3_deleteBucketMetadataTableConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected bucket owner of the general purpose bucket that you want to remove the metadata table configuration from.")
		s3_deleteBucketMetadataTableConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_deleteBucketMetadataTableConfigurationCmd)
}
