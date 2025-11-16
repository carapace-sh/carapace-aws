package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getRestoreJobMetadataCmd = &cobra.Command{
	Use:   "get-restore-job-metadata",
	Short: "This request returns the metadata for the specified restore job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getRestoreJobMetadataCmd).Standalone()

	backup_getRestoreJobMetadataCmd.Flags().String("restore-job-id", "", "This is a unique identifier of a restore job within Backup.")
	backup_getRestoreJobMetadataCmd.MarkFlagRequired("restore-job-id")
	backupCmd.AddCommand(backup_getRestoreJobMetadataCmd)
}
