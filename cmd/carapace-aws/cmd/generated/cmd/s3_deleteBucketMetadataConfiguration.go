package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteBucketMetadataConfigurationCmd = &cobra.Command{
	Use:   "delete-bucket-metadata-configuration",
	Short: "Deletes an S3 Metadata configuration from a general purpose bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteBucketMetadataConfigurationCmd).Standalone()

	s3_deleteBucketMetadataConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that you want to remove the metadata configuration from.")
	s3_deleteBucketMetadataConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected bucket owner of the general purpose bucket that you want to remove the metadata table configuration from.")
	s3_deleteBucketMetadataConfigurationCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_deleteBucketMetadataConfigurationCmd)
}
