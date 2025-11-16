package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_createPerformanceAnalysisReportCmd = &cobra.Command{
	Use:   "create-performance-analysis-report",
	Short: "Creates a new performance analysis report for a specific time period for the DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_createPerformanceAnalysisReportCmd).Standalone()

	pi_createPerformanceAnalysisReportCmd.Flags().String("end-time", "", "The end time defined for the analysis report.")
	pi_createPerformanceAnalysisReportCmd.Flags().String("identifier", "", "An immutable, Amazon Web Services Region-unique identifier for a data source.")
	pi_createPerformanceAnalysisReportCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights will return metrics.")
	pi_createPerformanceAnalysisReportCmd.Flags().String("start-time", "", "The start time defined for the analysis report.")
	pi_createPerformanceAnalysisReportCmd.Flags().String("tags", "", "The metadata assigned to the analysis report consisting of a key-value pair.")
	pi_createPerformanceAnalysisReportCmd.MarkFlagRequired("end-time")
	pi_createPerformanceAnalysisReportCmd.MarkFlagRequired("identifier")
	pi_createPerformanceAnalysisReportCmd.MarkFlagRequired("service-type")
	pi_createPerformanceAnalysisReportCmd.MarkFlagRequired("start-time")
	piCmd.AddCommand(pi_createPerformanceAnalysisReportCmd)
}
