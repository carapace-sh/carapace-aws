package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listReportJobsCmd = &cobra.Command{
	Use:   "list-report-jobs",
	Short: "Returns details about your report jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listReportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listReportJobsCmd).Standalone()

		backup_listReportJobsCmd.Flags().String("by-creation-after", "", "Returns only report jobs that were created after the date and time specified in Unix format and Coordinated Universal Time (UTC).")
		backup_listReportJobsCmd.Flags().String("by-creation-before", "", "Returns only report jobs that were created before the date and time specified in Unix format and Coordinated Universal Time (UTC).")
		backup_listReportJobsCmd.Flags().String("by-report-plan-name", "", "Returns only report jobs with the specified report plan name.")
		backup_listReportJobsCmd.Flags().String("by-status", "", "Returns only report jobs that are in the specified status.")
		backup_listReportJobsCmd.Flags().String("max-results", "", "The number of desired results from 1 to 1000. Optional.")
		backup_listReportJobsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	})
	backupCmd.AddCommand(backup_listReportJobsCmd)
}
