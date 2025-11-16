package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteBackupVaultNotificationsCmd = &cobra.Command{
	Use:   "delete-backup-vault-notifications",
	Short: "Deletes event notifications for the specified backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteBackupVaultNotificationsCmd).Standalone()

	backup_deleteBackupVaultNotificationsCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_deleteBackupVaultNotificationsCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_deleteBackupVaultNotificationsCmd)
}
