package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteReportGroupCmd = &cobra.Command{
	Use:   "delete-report-group",
	Short: "Deletes a report group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteReportGroupCmd).Standalone()

	codebuild_deleteReportGroupCmd.Flags().String("arn", "", "The ARN of the report group to delete.")
	codebuild_deleteReportGroupCmd.Flags().Bool("delete-reports", false, "If `true`, deletes any reports that belong to a report group before deleting the report group.")
	codebuild_deleteReportGroupCmd.Flags().Bool("no-delete-reports", false, "If `true`, deletes any reports that belong to a report group before deleting the report group.")
	codebuild_deleteReportGroupCmd.MarkFlagRequired("arn")
	codebuild_deleteReportGroupCmd.Flag("no-delete-reports").Hidden = true
	codebuildCmd.AddCommand(codebuild_deleteReportGroupCmd)
}
