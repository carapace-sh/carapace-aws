package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createBackupVaultCmd = &cobra.Command{
	Use:   "create-backup-vault",
	Short: "Creates a logical container where backups are stored.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createBackupVaultCmd).Standalone()

	backup_createBackupVaultCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_createBackupVaultCmd.Flags().String("backup-vault-tags", "", "The tags to assign to the backup vault.")
	backup_createBackupVaultCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.")
	backup_createBackupVaultCmd.Flags().String("encryption-key-arn", "", "The server-side encryption key that is used to protect your backups; for example, `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`.")
	backup_createBackupVaultCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_createBackupVaultCmd)
}
