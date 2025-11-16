package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_getReportGroupTrendCmd = &cobra.Command{
	Use:   "get-report-group-trend",
	Short: "Analyzes and accumulates test report values for the specified test reports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_getReportGroupTrendCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_getReportGroupTrendCmd).Standalone()

		codebuild_getReportGroupTrendCmd.Flags().String("num-of-reports", "", "The number of reports to analyze.")
		codebuild_getReportGroupTrendCmd.Flags().String("report-group-arn", "", "The ARN of the report group that contains the reports to analyze.")
		codebuild_getReportGroupTrendCmd.Flags().String("trend-field", "", "The test report value to accumulate.")
		codebuild_getReportGroupTrendCmd.MarkFlagRequired("report-group-arn")
		codebuild_getReportGroupTrendCmd.MarkFlagRequired("trend-field")
	})
	codebuildCmd.AddCommand(codebuild_getReportGroupTrendCmd)
}
