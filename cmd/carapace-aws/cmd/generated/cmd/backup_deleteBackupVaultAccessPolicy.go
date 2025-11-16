package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteBackupVaultAccessPolicyCmd = &cobra.Command{
	Use:   "delete-backup-vault-access-policy",
	Short: "Deletes the policy document that manages permissions on a backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteBackupVaultAccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_deleteBackupVaultAccessPolicyCmd).Standalone()

		backup_deleteBackupVaultAccessPolicyCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_deleteBackupVaultAccessPolicyCmd.MarkFlagRequired("backup-vault-name")
	})
	backupCmd.AddCommand(backup_deleteBackupVaultAccessPolicyCmd)
}
