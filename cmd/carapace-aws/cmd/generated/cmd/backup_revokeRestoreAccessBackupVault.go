package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_revokeRestoreAccessBackupVaultCmd = &cobra.Command{
	Use:   "revoke-restore-access-backup-vault",
	Short: "Revokes access to a restore access backup vault, removing the ability to restore from its recovery points and permanently deleting the vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_revokeRestoreAccessBackupVaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_revokeRestoreAccessBackupVaultCmd).Standalone()

		backup_revokeRestoreAccessBackupVaultCmd.Flags().String("backup-vault-name", "", "The name of the source backup vault associated with the restore access backup vault to be revoked.")
		backup_revokeRestoreAccessBackupVaultCmd.Flags().String("requester-comment", "", "A comment explaining the reason for revoking access to the restore access backup vault.")
		backup_revokeRestoreAccessBackupVaultCmd.Flags().String("restore-access-backup-vault-arn", "", "The ARN of the restore access backup vault to revoke.")
		backup_revokeRestoreAccessBackupVaultCmd.MarkFlagRequired("backup-vault-name")
		backup_revokeRestoreAccessBackupVaultCmd.MarkFlagRequired("restore-access-backup-vault-arn")
	})
	backupCmd.AddCommand(backup_revokeRestoreAccessBackupVaultCmd)
}
