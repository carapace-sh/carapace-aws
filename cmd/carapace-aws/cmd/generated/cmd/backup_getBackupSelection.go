package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getBackupSelectionCmd = &cobra.Command{
	Use:   "get-backup-selection",
	Short: "Returns selection metadata and a document in JSON format that specifies a list of resources that are associated with a backup plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getBackupSelectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getBackupSelectionCmd).Standalone()

		backup_getBackupSelectionCmd.Flags().String("backup-plan-id", "", "Uniquely identifies a backup plan.")
		backup_getBackupSelectionCmd.Flags().String("selection-id", "", "Uniquely identifies the body of a request to assign a set of resources to a backup plan.")
		backup_getBackupSelectionCmd.MarkFlagRequired("backup-plan-id")
		backup_getBackupSelectionCmd.MarkFlagRequired("selection-id")
	})
	backupCmd.AddCommand(backup_getBackupSelectionCmd)
}
