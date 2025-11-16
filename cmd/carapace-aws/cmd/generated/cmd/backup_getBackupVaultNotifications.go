package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getBackupVaultNotificationsCmd = &cobra.Command{
	Use:   "get-backup-vault-notifications",
	Short: "Returns event notifications for the specified backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getBackupVaultNotificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getBackupVaultNotificationsCmd).Standalone()

		backup_getBackupVaultNotificationsCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_getBackupVaultNotificationsCmd.MarkFlagRequired("backup-vault-name")
	})
	backupCmd.AddCommand(backup_getBackupVaultNotificationsCmd)
}
