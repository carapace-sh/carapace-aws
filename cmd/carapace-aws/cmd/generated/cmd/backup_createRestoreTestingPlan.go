package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createRestoreTestingPlanCmd = &cobra.Command{
	Use:   "create-restore-testing-plan",
	Short: "Creates a restore testing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createRestoreTestingPlanCmd).Standalone()

	backup_createRestoreTestingPlanCmd.Flags().String("creator-request-id", "", "This is a unique string that identifies the request and allows failed requests to be retriedwithout the risk of running the operation twice.")
	backup_createRestoreTestingPlanCmd.Flags().String("restore-testing-plan", "", "A restore testing plan must contain a unique `RestoreTestingPlanName` string you create and must contain a `ScheduleExpression` cron.")
	backup_createRestoreTestingPlanCmd.Flags().String("tags", "", "The tags to assign to the restore testing plan.")
	backup_createRestoreTestingPlanCmd.MarkFlagRequired("restore-testing-plan")
	backupCmd.AddCommand(backup_createRestoreTestingPlanCmd)
}
