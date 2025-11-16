package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateRestoreTestingSelectionCmd = &cobra.Command{
	Use:   "update-restore-testing-selection",
	Short: "Updates the specified restore testing selection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateRestoreTestingSelectionCmd).Standalone()

	backup_updateRestoreTestingSelectionCmd.Flags().String("restore-testing-plan-name", "", "The restore testing plan name is required to update the indicated testing plan.")
	backup_updateRestoreTestingSelectionCmd.Flags().String("restore-testing-selection", "", "To update your restore testing selection, you can use either protected resource ARNs or conditions, but not both.")
	backup_updateRestoreTestingSelectionCmd.Flags().String("restore-testing-selection-name", "", "The required restore testing selection name of the restore testing selection you wish to update.")
	backup_updateRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-plan-name")
	backup_updateRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-selection")
	backup_updateRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-selection-name")
	backupCmd.AddCommand(backup_updateRestoreTestingSelectionCmd)
}
