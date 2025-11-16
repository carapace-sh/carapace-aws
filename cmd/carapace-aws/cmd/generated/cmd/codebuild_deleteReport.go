package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteReportCmd = &cobra.Command{
	Use:   "delete-report",
	Short: "Deletes a report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_deleteReportCmd).Standalone()

		codebuild_deleteReportCmd.Flags().String("arn", "", "The ARN of the report to delete.")
		codebuild_deleteReportCmd.MarkFlagRequired("arn")
	})
	codebuildCmd.AddCommand(codebuild_deleteReportCmd)
}
