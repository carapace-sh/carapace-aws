package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeReportJobCmd = &cobra.Command{
	Use:   "describe-report-job",
	Short: "Returns the details associated with creating a report as specified by its `ReportJobId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeReportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_describeReportJobCmd).Standalone()

		backup_describeReportJobCmd.Flags().String("report-job-id", "", "The identifier of the report job.")
		backup_describeReportJobCmd.MarkFlagRequired("report-job-id")
	})
	backupCmd.AddCommand(backup_describeReportJobCmd)
}
