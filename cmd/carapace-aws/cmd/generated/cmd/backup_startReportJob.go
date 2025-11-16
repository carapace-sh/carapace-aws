package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_startReportJobCmd = &cobra.Command{
	Use:   "start-report-job",
	Short: "Starts an on-demand report job for the specified report plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_startReportJobCmd).Standalone()

	backup_startReportJobCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `StartReportJobInput`.")
	backup_startReportJobCmd.Flags().String("report-plan-name", "", "The unique name of a report plan.")
	backup_startReportJobCmd.MarkFlagRequired("report-plan-name")
	backupCmd.AddCommand(backup_startReportJobCmd)
}
