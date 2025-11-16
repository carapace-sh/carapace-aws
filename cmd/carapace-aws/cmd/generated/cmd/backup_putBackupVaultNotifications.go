package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_putBackupVaultNotificationsCmd = &cobra.Command{
	Use:   "put-backup-vault-notifications",
	Short: "Turns on notifications on a backup vault for the specified topic and events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_putBackupVaultNotificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_putBackupVaultNotificationsCmd).Standalone()

		backup_putBackupVaultNotificationsCmd.Flags().String("backup-vault-events", "", "An array of events that indicate the status of jobs to back up resources to the backup vault.")
		backup_putBackupVaultNotificationsCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_putBackupVaultNotificationsCmd.Flags().String("snstopic-arn", "", "The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events; for example, `arn:aws:sns:us-west-2:111122223333:MyVaultTopic`.")
		backup_putBackupVaultNotificationsCmd.MarkFlagRequired("backup-vault-events")
		backup_putBackupVaultNotificationsCmd.MarkFlagRequired("backup-vault-name")
		backup_putBackupVaultNotificationsCmd.MarkFlagRequired("snstopic-arn")
	})
	backupCmd.AddCommand(backup_putBackupVaultNotificationsCmd)
}
