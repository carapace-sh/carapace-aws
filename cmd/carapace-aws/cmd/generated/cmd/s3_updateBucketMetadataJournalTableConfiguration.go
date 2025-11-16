package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_updateBucketMetadataJournalTableConfigurationCmd = &cobra.Command{
	Use:   "update-bucket-metadata-journal-table-configuration",
	Short: "Enables or disables journal table record expiration for an S3 Metadata configuration on a general purpose bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_updateBucketMetadataJournalTableConfigurationCmd).Standalone()

	s3_updateBucketMetadataJournalTableConfigurationCmd.Flags().String("bucket", "", "The general purpose bucket that corresponds to the metadata configuration that you want to enable or disable journal table record expiration for.")
	s3_updateBucketMetadataJournalTableConfigurationCmd.Flags().String("checksum-algorithm", "", "The checksum algorithm to use with your journal table configuration.")
	s3_updateBucketMetadataJournalTableConfigurationCmd.Flags().String("content-md5", "", "The `Content-MD5` header for the journal table configuration.")
	s3_updateBucketMetadataJournalTableConfigurationCmd.Flags().String("expected-bucket-owner", "", "The expected owner of the general purpose bucket that corresponds to the metadata table configuration that you want to enable or disable journal table record expiration for.")
	s3_updateBucketMetadataJournalTableConfigurationCmd.Flags().String("journal-table-configuration", "", "The contents of your journal table configuration.")
	s3_updateBucketMetadataJournalTableConfigurationCmd.MarkFlagRequired("bucket")
	s3_updateBucketMetadataJournalTableConfigurationCmd.MarkFlagRequired("journal-table-configuration")
	s3Cmd.AddCommand(s3_updateBucketMetadataJournalTableConfigurationCmd)
}
