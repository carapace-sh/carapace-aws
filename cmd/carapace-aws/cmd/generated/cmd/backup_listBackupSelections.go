package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listBackupSelectionsCmd = &cobra.Command{
	Use:   "list-backup-selections",
	Short: "Returns an array containing metadata of the resources associated with the target backup plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listBackupSelectionsCmd).Standalone()

	backup_listBackupSelectionsCmd.Flags().String("backup-plan-id", "", "Uniquely identifies a backup plan.")
	backup_listBackupSelectionsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listBackupSelectionsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backup_listBackupSelectionsCmd.MarkFlagRequired("backup-plan-id")
	backupCmd.AddCommand(backup_listBackupSelectionsCmd)
}
