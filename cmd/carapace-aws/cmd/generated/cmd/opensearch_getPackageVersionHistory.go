package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getPackageVersionHistoryCmd = &cobra.Command{
	Use:   "get-package-version-history",
	Short: "Returns a list of Amazon OpenSearch Service package versions, along with their creation time, commit message, and plugin properties (if the package is a zip plugin package).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getPackageVersionHistoryCmd).Standalone()

	opensearch_getPackageVersionHistoryCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_getPackageVersionHistoryCmd.Flags().String("next-token", "", "If your initial `GetPackageVersionHistory` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `GetPackageVersionHistory` operations, which returns results in the next page.")
	opensearch_getPackageVersionHistoryCmd.Flags().String("package-id", "", "The unique identifier of the package.")
	opensearch_getPackageVersionHistoryCmd.MarkFlagRequired("package-id")
	opensearchCmd.AddCommand(opensearch_getPackageVersionHistoryCmd)
}
