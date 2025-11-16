package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateRestoreTestingPlanCmd = &cobra.Command{
	Use:   "update-restore-testing-plan",
	Short: "This request will send changes to your specified restore testing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateRestoreTestingPlanCmd).Standalone()

	backup_updateRestoreTestingPlanCmd.Flags().String("restore-testing-plan", "", "Specifies the body of a restore testing plan.")
	backup_updateRestoreTestingPlanCmd.Flags().String("restore-testing-plan-name", "", "The name of the restore testing plan name.")
	backup_updateRestoreTestingPlanCmd.MarkFlagRequired("restore-testing-plan")
	backup_updateRestoreTestingPlanCmd.MarkFlagRequired("restore-testing-plan-name")
	backupCmd.AddCommand(backup_updateRestoreTestingPlanCmd)
}
