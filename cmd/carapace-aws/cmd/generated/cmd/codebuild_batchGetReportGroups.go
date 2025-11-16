package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetReportGroupsCmd = &cobra.Command{
	Use:   "batch-get-report-groups",
	Short: "Returns an array of report groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetReportGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_batchGetReportGroupsCmd).Standalone()

		codebuild_batchGetReportGroupsCmd.Flags().String("report-group-arns", "", "An array of report group ARNs that identify the report groups to return.")
		codebuild_batchGetReportGroupsCmd.MarkFlagRequired("report-group-arns")
	})
	codebuildCmd.AddCommand(codebuild_batchGetReportGroupsCmd)
}
