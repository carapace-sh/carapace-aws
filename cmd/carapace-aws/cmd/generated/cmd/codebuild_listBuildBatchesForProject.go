package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listBuildBatchesForProjectCmd = &cobra.Command{
	Use:   "list-build-batches-for-project",
	Short: "Retrieves the identifiers of the build batches for a specific project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listBuildBatchesForProjectCmd).Standalone()

	codebuild_listBuildBatchesForProjectCmd.Flags().String("filter", "", "A `BuildBatchFilter` object that specifies the filters for the search.")
	codebuild_listBuildBatchesForProjectCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	codebuild_listBuildBatchesForProjectCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous call to `ListBuildBatchesForProject`.")
	codebuild_listBuildBatchesForProjectCmd.Flags().String("project-name", "", "The name of the project.")
	codebuild_listBuildBatchesForProjectCmd.Flags().String("sort-order", "", "Specifies the sort order of the returned items.")
	codebuildCmd.AddCommand(codebuild_listBuildBatchesForProjectCmd)
}
