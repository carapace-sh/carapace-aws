package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createRestoreTestingSelectionCmd = &cobra.Command{
	Use:   "create-restore-testing-selection",
	Short: "This request can be sent after CreateRestoreTestingPlan request returns successfully.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createRestoreTestingSelectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_createRestoreTestingSelectionCmd).Standalone()

		backup_createRestoreTestingSelectionCmd.Flags().String("creator-request-id", "", "This is an optional unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.")
		backup_createRestoreTestingSelectionCmd.Flags().String("restore-testing-plan-name", "", "Input the restore testing plan name that was returned from the related CreateRestoreTestingPlan request.")
		backup_createRestoreTestingSelectionCmd.Flags().String("restore-testing-selection", "", "This consists of `RestoreTestingSelectionName`, `ProtectedResourceType`, and one of the following:")
		backup_createRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-plan-name")
		backup_createRestoreTestingSelectionCmd.MarkFlagRequired("restore-testing-selection")
	})
	backupCmd.AddCommand(backup_createRestoreTestingSelectionCmd)
}
