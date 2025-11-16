package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getRestoreTestingPlanCmd = &cobra.Command{
	Use:   "get-restore-testing-plan",
	Short: "Returns `RestoreTestingPlan` details for the specified `RestoreTestingPlanName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getRestoreTestingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getRestoreTestingPlanCmd).Standalone()

		backup_getRestoreTestingPlanCmd.Flags().String("restore-testing-plan-name", "", "Required unique name of the restore testing plan.")
		backup_getRestoreTestingPlanCmd.MarkFlagRequired("restore-testing-plan-name")
	})
	backupCmd.AddCommand(backup_getRestoreTestingPlanCmd)
}
