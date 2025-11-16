package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketMetadataConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-metadata-configuration",
	Short: "Retrieves the S3 Metadata configuration for a general purpose bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketMetadataConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketMetadataConfigurationCmd).Standalone()

		s3_getBucketMetadataConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that corresponds to the metadata configuration that you want to retrieve.")
		s3_getBucketMetadataConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected owner of the general purpose bucket that you want to retrieve the metadata table configuration for.")
		s3_getBucketMetadataConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_getBucketMetadataConfigurationCmd)
}
