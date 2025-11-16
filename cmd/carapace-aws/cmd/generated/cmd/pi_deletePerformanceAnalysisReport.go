package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_deletePerformanceAnalysisReportCmd = &cobra.Command{
	Use:   "delete-performance-analysis-report",
	Short: "Deletes a performance analysis report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_deletePerformanceAnalysisReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_deletePerformanceAnalysisReportCmd).Standalone()

		pi_deletePerformanceAnalysisReportCmd.Flags().String("analysis-report-id", "", "The unique identifier of the analysis report for deletion.")
		pi_deletePerformanceAnalysisReportCmd.Flags().String("identifier", "", "An immutable identifier for a data source that is unique for an Amazon Web Services Region.")
		pi_deletePerformanceAnalysisReportCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights will return metrics.")
		pi_deletePerformanceAnalysisReportCmd.MarkFlagRequired("analysis-report-id")
		pi_deletePerformanceAnalysisReportCmd.MarkFlagRequired("identifier")
		pi_deletePerformanceAnalysisReportCmd.MarkFlagRequired("service-type")
	})
	piCmd.AddCommand(pi_deletePerformanceAnalysisReportCmd)
}
