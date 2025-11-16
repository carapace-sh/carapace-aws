package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteBackupVaultLockConfigurationCmd = &cobra.Command{
	Use:   "delete-backup-vault-lock-configuration",
	Short: "Deletes Backup Vault Lock from a backup vault specified by a backup vault name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteBackupVaultLockConfigurationCmd).Standalone()

	backup_deleteBackupVaultLockConfigurationCmd.Flags().String("backup-vault-name", "", "The name of the backup vault from which to delete Backup Vault Lock.")
	backup_deleteBackupVaultLockConfigurationCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_deleteBackupVaultLockConfigurationCmd)
}
