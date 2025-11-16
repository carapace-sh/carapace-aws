package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_retryBuildCmd = &cobra.Command{
	Use:   "retry-build",
	Short: "Restarts a build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_retryBuildCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_retryBuildCmd).Standalone()

		codebuild_retryBuildCmd.Flags().String("id", "", "Specifies the identifier of the build to restart.")
		codebuild_retryBuildCmd.Flags().String("idempotency-token", "", "A unique, case sensitive identifier you provide to ensure the idempotency of the `RetryBuild` request.")
	})
	codebuildCmd.AddCommand(codebuild_retryBuildCmd)
}
