package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listBuildBatchesCmd = &cobra.Command{
	Use:   "list-build-batches",
	Short: "Retrieves the identifiers of your build batches in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listBuildBatchesCmd).Standalone()

	codebuild_listBuildBatchesCmd.Flags().String("filter", "", "A `BuildBatchFilter` object that specifies the filters for the search.")
	codebuild_listBuildBatchesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	codebuild_listBuildBatchesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous call to `ListBuildBatches`.")
	codebuild_listBuildBatchesCmd.Flags().String("sort-order", "", "Specifies the sort order of the returned items.")
	codebuildCmd.AddCommand(codebuild_listBuildBatchesCmd)
}
