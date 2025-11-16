package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_stopBuildBatchCmd = &cobra.Command{
	Use:   "stop-build-batch",
	Short: "Stops a running batch build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_stopBuildBatchCmd).Standalone()

	codebuild_stopBuildBatchCmd.Flags().String("id", "", "The identifier of the batch build to stop.")
	codebuild_stopBuildBatchCmd.MarkFlagRequired("id")
	codebuildCmd.AddCommand(codebuild_stopBuildBatchCmd)
}
