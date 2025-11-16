package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getBackupVaultAccessPolicyCmd = &cobra.Command{
	Use:   "get-backup-vault-access-policy",
	Short: "Returns the access policy document that is associated with the named backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getBackupVaultAccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getBackupVaultAccessPolicyCmd).Standalone()

		backup_getBackupVaultAccessPolicyCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_getBackupVaultAccessPolicyCmd.MarkFlagRequired("backup-vault-name")
	})
	backupCmd.AddCommand(backup_getBackupVaultAccessPolicyCmd)
}
