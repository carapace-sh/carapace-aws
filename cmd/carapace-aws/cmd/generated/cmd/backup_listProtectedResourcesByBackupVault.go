package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listProtectedResourcesByBackupVaultCmd = &cobra.Command{
	Use:   "list-protected-resources-by-backup-vault",
	Short: "This request lists the protected resources corresponding to each backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listProtectedResourcesByBackupVaultCmd).Standalone()

	backup_listProtectedResourcesByBackupVaultCmd.Flags().String("backup-vault-account-id", "", "The list of protected resources by backup vault within the vault(s) you specify by account ID.")
	backup_listProtectedResourcesByBackupVaultCmd.Flags().String("backup-vault-name", "", "The list of protected resources by backup vault within the vault(s) you specify by name.")
	backup_listProtectedResourcesByBackupVaultCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listProtectedResourcesByBackupVaultCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backup_listProtectedResourcesByBackupVaultCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_listProtectedResourcesByBackupVaultCmd)
}
