package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_updateReportGroupCmd = &cobra.Command{
	Use:   "update-report-group",
	Short: "Updates a report group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_updateReportGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_updateReportGroupCmd).Standalone()

		codebuild_updateReportGroupCmd.Flags().String("arn", "", "The ARN of the report group to update.")
		codebuild_updateReportGroupCmd.Flags().String("export-config", "", "Used to specify an updated export type.")
		codebuild_updateReportGroupCmd.Flags().String("tags", "", "An updated list of tag key and value pairs associated with this report group.")
		codebuild_updateReportGroupCmd.MarkFlagRequired("arn")
	})
	codebuildCmd.AddCommand(codebuild_updateReportGroupCmd)
}
