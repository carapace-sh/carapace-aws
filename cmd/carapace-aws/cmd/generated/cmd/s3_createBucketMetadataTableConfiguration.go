package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_createBucketMetadataTableConfigurationCmd = &cobra.Command{
	Use:   "create-bucket-metadata-table-configuration",
	Short: "We recommend that you create your S3 Metadata configurations by using the V2 [CreateBucketMetadataConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html) API operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_createBucketMetadataTableConfigurationCmd).Standalone()

	s3_createBucketMetadataTableConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that you want to create the metadata table configuration for.")
	s3_createBucketMetadataTableConfigurationCmd.Flags().String("checksum-algorithm", "", "The checksum algorithm to use with your metadata table configuration.")
	s3_createBucketMetadataTableConfigurationCmd.Flags().String("content-md5", "", "The `Content-MD5` header for the metadata table configuration.")
	s3_createBucketMetadataTableConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected owner of the general purpose bucket that corresponds to your metadata table configuration.")
	s3_createBucketMetadataTableConfigurationCmd.Flags().String("metadata-table-configuration", "", "The contents of your metadata table configuration.")
	s3_createBucketMetadataTableConfigurationCmd.MarkFlagRequired("bucket")
	s3_createBucketMetadataTableConfigurationCmd.MarkFlagRequired("metadata-table-configuration")
	s3Cmd.AddCommand(s3_createBucketMetadataTableConfigurationCmd)
}
