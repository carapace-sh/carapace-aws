package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeReportPlanCmd = &cobra.Command{
	Use:   "describe-report-plan",
	Short: "Returns a list of all report plans for an Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeReportPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_describeReportPlanCmd).Standalone()

		backup_describeReportPlanCmd.Flags().String("report-plan-name", "", "The unique name of a report plan.")
		backup_describeReportPlanCmd.MarkFlagRequired("report-plan-name")
	})
	backupCmd.AddCommand(backup_describeReportPlanCmd)
}
