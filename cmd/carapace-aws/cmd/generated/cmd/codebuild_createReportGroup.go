package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_createReportGroupCmd = &cobra.Command{
	Use:   "create-report-group",
	Short: "Creates a report group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_createReportGroupCmd).Standalone()

	codebuild_createReportGroupCmd.Flags().String("export-config", "", "A `ReportExportConfig` object that contains information about where the report group test results are exported.")
	codebuild_createReportGroupCmd.Flags().String("name", "", "The name of the report group.")
	codebuild_createReportGroupCmd.Flags().String("tags", "", "A list of tag key and value pairs associated with this report group.")
	codebuild_createReportGroupCmd.Flags().String("type", "", "The type of report group.")
	codebuild_createReportGroupCmd.MarkFlagRequired("export-config")
	codebuild_createReportGroupCmd.MarkFlagRequired("name")
	codebuild_createReportGroupCmd.MarkFlagRequired("type")
	codebuildCmd.AddCommand(codebuild_createReportGroupCmd)
}
