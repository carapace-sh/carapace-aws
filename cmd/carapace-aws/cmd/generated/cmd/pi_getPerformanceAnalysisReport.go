package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_getPerformanceAnalysisReportCmd = &cobra.Command{
	Use:   "get-performance-analysis-report",
	Short: "Retrieves the report including the report ID, status, time details, and the insights with recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_getPerformanceAnalysisReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_getPerformanceAnalysisReportCmd).Standalone()

		pi_getPerformanceAnalysisReportCmd.Flags().String("accept-language", "", "The text language in the report.")
		pi_getPerformanceAnalysisReportCmd.Flags().String("analysis-report-id", "", "A unique identifier of the created analysis report.")
		pi_getPerformanceAnalysisReportCmd.Flags().String("identifier", "", "An immutable identifier for a data source that is unique for an Amazon Web Services Region.")
		pi_getPerformanceAnalysisReportCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights will return metrics.")
		pi_getPerformanceAnalysisReportCmd.Flags().String("text-format", "", "Indicates the text format in the report.")
		pi_getPerformanceAnalysisReportCmd.MarkFlagRequired("analysis-report-id")
		pi_getPerformanceAnalysisReportCmd.MarkFlagRequired("identifier")
		pi_getPerformanceAnalysisReportCmd.MarkFlagRequired("service-type")
	})
	piCmd.AddCommand(pi_getPerformanceAnalysisReportCmd)
}
