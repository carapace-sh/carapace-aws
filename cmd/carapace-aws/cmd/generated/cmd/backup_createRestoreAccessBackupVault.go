package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createRestoreAccessBackupVaultCmd = &cobra.Command{
	Use:   "create-restore-access-backup-vault",
	Short: "Creates a restore access backup vault that provides temporary access to recovery points in a logically air-gapped backup vault, subject to MPA approval.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createRestoreAccessBackupVaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_createRestoreAccessBackupVaultCmd).Standalone()

		backup_createRestoreAccessBackupVaultCmd.Flags().String("backup-vault-name", "", "The name of the backup vault to associate with an MPA approval team.")
		backup_createRestoreAccessBackupVaultCmd.Flags().String("backup-vault-tags", "", "Optional tags to assign to the restore access backup vault.")
		backup_createRestoreAccessBackupVaultCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and allows failed requests to be retried without the risk of executing the operation twice.")
		backup_createRestoreAccessBackupVaultCmd.Flags().String("requester-comment", "", "A comment explaining the reason for requesting restore access to the backup vault.")
		backup_createRestoreAccessBackupVaultCmd.Flags().String("source-backup-vault-arn", "", "The ARN of the source backup vault containing the recovery points to which temporary access is requested.")
		backup_createRestoreAccessBackupVaultCmd.MarkFlagRequired("source-backup-vault-arn")
	})
	backupCmd.AddCommand(backup_createRestoreAccessBackupVaultCmd)
}
