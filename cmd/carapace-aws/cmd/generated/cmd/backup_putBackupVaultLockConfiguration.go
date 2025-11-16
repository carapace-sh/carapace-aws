package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_putBackupVaultLockConfigurationCmd = &cobra.Command{
	Use:   "put-backup-vault-lock-configuration",
	Short: "Applies Backup Vault Lock to a backup vault, preventing attempts to delete any recovery point stored in or created in a backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_putBackupVaultLockConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_putBackupVaultLockConfigurationCmd).Standalone()

		backup_putBackupVaultLockConfigurationCmd.Flags().String("backup-vault-name", "", "The Backup Vault Lock configuration that specifies the name of the backup vault it protects.")
		backup_putBackupVaultLockConfigurationCmd.Flags().String("changeable-for-days", "", "The Backup Vault Lock configuration that specifies the number of days before the lock date.")
		backup_putBackupVaultLockConfigurationCmd.Flags().String("max-retention-days", "", "The Backup Vault Lock configuration that specifies the maximum retention period that the vault retains its recovery points.")
		backup_putBackupVaultLockConfigurationCmd.Flags().String("min-retention-days", "", "The Backup Vault Lock configuration that specifies the minimum retention period that the vault retains its recovery points.")
		backup_putBackupVaultLockConfigurationCmd.MarkFlagRequired("backup-vault-name")
	})
	backupCmd.AddCommand(backup_putBackupVaultLockConfigurationCmd)
}
