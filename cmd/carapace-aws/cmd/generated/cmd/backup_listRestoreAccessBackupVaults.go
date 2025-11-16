package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRestoreAccessBackupVaultsCmd = &cobra.Command{
	Use:   "list-restore-access-backup-vaults",
	Short: "Returns a list of restore access backup vaults associated with a specified backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRestoreAccessBackupVaultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listRestoreAccessBackupVaultsCmd).Standalone()

		backup_listRestoreAccessBackupVaultsCmd.Flags().String("backup-vault-name", "", "The name of the backup vault for which to list associated restore access backup vaults.")
		backup_listRestoreAccessBackupVaultsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		backup_listRestoreAccessBackupVaultsCmd.Flags().String("next-token", "", "The pagination token from a previous request to retrieve the next set of results.")
		backup_listRestoreAccessBackupVaultsCmd.MarkFlagRequired("backup-vault-name")
	})
	backupCmd.AddCommand(backup_listRestoreAccessBackupVaultsCmd)
}
