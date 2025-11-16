package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listBackupVaultsCmd = &cobra.Command{
	Use:   "list-backup-vaults",
	Short: "Returns a list of recovery point storage containers along with information about them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listBackupVaultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listBackupVaultsCmd).Standalone()

		backup_listBackupVaultsCmd.Flags().String("by-shared", "", "This parameter will sort the list of vaults by shared vaults.")
		backup_listBackupVaultsCmd.Flags().String("by-vault-type", "", "This parameter will sort the list of vaults by vault type.")
		backup_listBackupVaultsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
		backup_listBackupVaultsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	})
	backupCmd.AddCommand(backup_listBackupVaultsCmd)
}
