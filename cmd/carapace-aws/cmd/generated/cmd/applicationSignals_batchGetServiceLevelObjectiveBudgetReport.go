package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_batchGetServiceLevelObjectiveBudgetReportCmd = &cobra.Command{
	Use:   "batch-get-service-level-objective-budget-report",
	Short: "Use this operation to retrieve one or more *service level objective (SLO) budget reports*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_batchGetServiceLevelObjectiveBudgetReportCmd).Standalone()

	applicationSignals_batchGetServiceLevelObjectiveBudgetReportCmd.Flags().String("slo-ids", "", "An array containing the IDs of the service level objectives that you want to include in the report.")
	applicationSignals_batchGetServiceLevelObjectiveBudgetReportCmd.Flags().String("timestamp", "", "The date and time that you want the report to be for.")
	applicationSignals_batchGetServiceLevelObjectiveBudgetReportCmd.MarkFlagRequired("slo-ids")
	applicationSignals_batchGetServiceLevelObjectiveBudgetReportCmd.MarkFlagRequired("timestamp")
	applicationSignalsCmd.AddCommand(applicationSignals_batchGetServiceLevelObjectiveBudgetReportCmd)
}
