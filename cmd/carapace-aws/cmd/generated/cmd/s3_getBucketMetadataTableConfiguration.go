package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketMetadataTableConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-metadata-table-configuration",
	Short: "We recommend that you retrieve your S3 Metadata configurations by using the V2 [GetBucketMetadataTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataTableConfiguration.html) API operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketMetadataTableConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketMetadataTableConfigurationCmd).Standalone()

		s3_getBucketMetadataTableConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that corresponds to the metadata table configuration that you want to retrieve.")
		s3_getBucketMetadataTableConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected owner of the general purpose bucket that you want to retrieve the metadata table configuration for.")
		s3_getBucketMetadataTableConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketMetadataTableConfigurationCmd)
}
