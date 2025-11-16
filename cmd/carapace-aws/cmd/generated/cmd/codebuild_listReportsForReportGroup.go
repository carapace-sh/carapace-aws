package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listReportsForReportGroupCmd = &cobra.Command{
	Use:   "list-reports-for-report-group",
	Short: "Returns a list of ARNs for the reports that belong to a `ReportGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listReportsForReportGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listReportsForReportGroupCmd).Standalone()

		codebuild_listReportsForReportGroupCmd.Flags().String("filter", "", "A `ReportFilter` object used to filter the returned reports.")
		codebuild_listReportsForReportGroupCmd.Flags().String("max-results", "", "The maximum number of paginated reports in this report group returned per response.")
		codebuild_listReportsForReportGroupCmd.Flags().String("next-token", "", "During a previous call, the maximum number of items that can be returned is the value specified in `maxResults`.")
		codebuild_listReportsForReportGroupCmd.Flags().String("report-group-arn", "", "The ARN of the report group for which you want to return report ARNs.")
		codebuild_listReportsForReportGroupCmd.Flags().String("sort-order", "", "Use to specify whether the results are returned in ascending or descending order.")
		codebuild_listReportsForReportGroupCmd.MarkFlagRequired("report-group-arn")
	})
	codebuildCmd.AddCommand(codebuild_listReportsForReportGroupCmd)
}
