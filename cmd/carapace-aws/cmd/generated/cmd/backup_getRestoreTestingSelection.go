package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getRestoreTestingSelectionCmd = &cobra.Command{
	Use:   "get-restore-testing-selection",
	Short: "Returns RestoreTestingSelection, which displays resources and elements of the restore testing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getRestoreTestingSelectionCmd).Standalone()

	backup_getRestoreTestingSelectionCmd.Flags().String("restore-testing-plan-name", "", "Required unique name of the restore testing plan.")
	backup_getRestoreTestingSelectionCmd.Flags().String("restore-testing-selection-name", "", "Required unique name of the restore testing selection.")
	backup_getRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-plan-name")
	backup_getRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-selection-name")
	backupCmd.AddCommand(backup_getRestoreTestingSelectionCmd)
}
