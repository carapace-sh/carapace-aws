package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRecoveryPointsByBackupVaultCmd = &cobra.Command{
	Use:   "list-recovery-points-by-backup-vault",
	Short: "Returns detailed information about the recovery points stored in a backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRecoveryPointsByBackupVaultCmd).Standalone()

	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("backup-vault-account-id", "", "This parameter will sort the list of recovery points by account ID.")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("by-backup-plan-id", "", "Returns only recovery points that match the specified backup plan ID.")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("by-created-after", "", "Returns only recovery points that were created after the specified timestamp.")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("by-created-before", "", "Returns only recovery points that were created before the specified timestamp.")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("by-parent-recovery-point-arn", "", "This returns only recovery points that match the specified parent (composite) recovery point Amazon Resource Name (ARN).")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("by-resource-arn", "", "Returns only recovery points that match the specified resource Amazon Resource Name (ARN).")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("by-resource-type", "", "Returns only recovery points that match the specified resource type(s):")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listRecoveryPointsByBackupVaultCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backup_listRecoveryPointsByBackupVaultCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_listRecoveryPointsByBackupVaultCmd)
}
