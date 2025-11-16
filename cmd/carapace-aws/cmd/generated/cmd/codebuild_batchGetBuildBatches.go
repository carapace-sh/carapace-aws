package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetBuildBatchesCmd = &cobra.Command{
	Use:   "batch-get-build-batches",
	Short: "Retrieves information about one or more batch builds.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetBuildBatchesCmd).Standalone()

	codebuild_batchGetBuildBatchesCmd.Flags().String("ids", "", "An array that contains the batch build identifiers to retrieve.")
	codebuild_batchGetBuildBatchesCmd.MarkFlagRequired("ids")
	codebuildCmd.AddCommand(codebuild_batchGetBuildBatchesCmd)
}
