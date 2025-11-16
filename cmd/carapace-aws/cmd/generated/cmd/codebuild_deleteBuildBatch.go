package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteBuildBatchCmd = &cobra.Command{
	Use:   "delete-build-batch",
	Short: "Deletes a batch build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteBuildBatchCmd).Standalone()

	codebuild_deleteBuildBatchCmd.Flags().String("id", "", "The identifier of the batch build to delete.")
	codebuild_deleteBuildBatchCmd.MarkFlagRequired("id")
	codebuildCmd.AddCommand(codebuild_deleteBuildBatchCmd)
}
