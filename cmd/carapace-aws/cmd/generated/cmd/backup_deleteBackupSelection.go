package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteBackupSelectionCmd = &cobra.Command{
	Use:   "delete-backup-selection",
	Short: "Deletes the resource selection associated with a backup plan that is specified by the `SelectionId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteBackupSelectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_deleteBackupSelectionCmd).Standalone()

		backup_deleteBackupSelectionCmd.Flags().String("backup-plan-id", "", "Uniquely identifies a backup plan.")
		backup_deleteBackupSelectionCmd.Flags().String("selection-id", "", "Uniquely identifies the body of a request to assign a set of resources to a backup plan.")
		backup_deleteBackupSelectionCmd.MarkFlagRequired("backup-plan-id")
		backup_deleteBackupSelectionCmd.MarkFlagRequired("selection-id")
	})
	backupCmd.AddCommand(backup_deleteBackupSelectionCmd)
}
