package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_retryBuildBatchCmd = &cobra.Command{
	Use:   "retry-build-batch",
	Short: "Restarts a failed batch build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_retryBuildBatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_retryBuildBatchCmd).Standalone()

		codebuild_retryBuildBatchCmd.Flags().String("id", "", "Specifies the identifier of the batch build to restart.")
		codebuild_retryBuildBatchCmd.Flags().String("idempotency-token", "", "A unique, case sensitive identifier you provide to ensure the idempotency of the `RetryBuildBatch` request.")
		codebuild_retryBuildBatchCmd.Flags().String("retry-type", "", "Specifies the type of retry to perform.")
	})
	codebuildCmd.AddCommand(codebuild_retryBuildBatchCmd)
}
