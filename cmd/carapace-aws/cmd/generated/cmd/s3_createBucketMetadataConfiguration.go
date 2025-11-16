package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_createBucketMetadataConfigurationCmd = &cobra.Command{
	Use:   "create-bucket-metadata-configuration",
	Short: "Creates an S3 Metadata V2 metadata configuration for a general purpose bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_createBucketMetadataConfigurationCmd).Standalone()

	s3_createBucketMetadataConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that you want to create the metadata configuration for.")
	s3_createBucketMetadataConfigurationCmd.Flags().String("checksum-algorithm", "", "The checksum algorithm to use with your metadata configuration.")
	s3_createBucketMetadataConfigurationCmd.Flags().String("content-md5", "", "The `Content-MD5` header for the metadata configuration.")
	s3_createBucketMetadataConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected owner of the general purpose bucket that corresponds to your metadata configuration.")
	s3_createBucketMetadataConfigurationCmd.Flags().String("metadata-configuration", "", "The contents of your metadata configuration.")
	s3_createBucketMetadataConfigurationCmd.MarkFlagRequired("bucket")
	s3_createBucketMetadataConfigurationCmd.MarkFlagRequired("metadata-configuration")
	s3Cmd.AddCommand(s3_createBucketMetadataConfigurationCmd)
}
