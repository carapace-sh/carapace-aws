package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getRecoveryPointRestoreMetadataCmd = &cobra.Command{
	Use:   "get-recovery-point-restore-metadata",
	Short: "Returns a set of metadata key-value pairs that were used to create the backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getRecoveryPointRestoreMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getRecoveryPointRestoreMetadataCmd).Standalone()

		backup_getRecoveryPointRestoreMetadataCmd.Flags().String("backup-vault-account-id", "", "The account ID of the specified backup vault.")
		backup_getRecoveryPointRestoreMetadataCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_getRecoveryPointRestoreMetadataCmd.Flags().String("recovery-point-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
		backup_getRecoveryPointRestoreMetadataCmd.MarkFlagRequired("backup-vault-name")
		backup_getRecoveryPointRestoreMetadataCmd.MarkFlagRequired("recovery-point-arn")
	})
	backupCmd.AddCommand(backup_getRecoveryPointRestoreMetadataCmd)
}
