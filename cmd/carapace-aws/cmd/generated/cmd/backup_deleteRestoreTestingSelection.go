package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteRestoreTestingSelectionCmd = &cobra.Command{
	Use:   "delete-restore-testing-selection",
	Short: "Input the Restore Testing Plan name and Restore Testing Selection name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteRestoreTestingSelectionCmd).Standalone()

	backup_deleteRestoreTestingSelectionCmd.Flags().String("restore-testing-plan-name", "", "Required unique name of the restore testing plan that contains the restore testing selection you wish to delete.")
	backup_deleteRestoreTestingSelectionCmd.Flags().String("restore-testing-selection-name", "", "Required unique name of the restore testing selection you wish to delete.")
	backup_deleteRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-plan-name")
	backup_deleteRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-selection-name")
	backupCmd.AddCommand(backup_deleteRestoreTestingSelectionCmd)
}
