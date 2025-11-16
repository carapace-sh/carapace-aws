package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetReportsCmd = &cobra.Command{
	Use:   "batch-get-reports",
	Short: "Returns an array of reports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_batchGetReportsCmd).Standalone()

		codebuild_batchGetReportsCmd.Flags().String("report-arns", "", "An array of ARNs that identify the `Report` objects to return.")
		codebuild_batchGetReportsCmd.MarkFlagRequired("report-arns")
	})
	codebuildCmd.AddCommand(codebuild_batchGetReportsCmd)
}
