package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteReportPlanCmd = &cobra.Command{
	Use:   "delete-report-plan",
	Short: "Deletes the report plan specified by a report plan name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteReportPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_deleteReportPlanCmd).Standalone()

		backup_deleteReportPlanCmd.Flags().String("report-plan-name", "", "The unique name of a report plan.")
		backup_deleteReportPlanCmd.MarkFlagRequired("report-plan-name")
	})
	backupCmd.AddCommand(backup_deleteReportPlanCmd)
}
