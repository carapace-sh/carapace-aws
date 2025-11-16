package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describePackagesCmd = &cobra.Command{
	Use:   "describe-packages",
	Short: "Describes all packages available to OpenSearch Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describePackagesCmd).Standalone()

	opensearch_describePackagesCmd.Flags().String("filters", "", "Only returns packages that match the `DescribePackagesFilterList` values.")
	opensearch_describePackagesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_describePackagesCmd.Flags().String("next-token", "", "If your initial `DescribePackageFilters` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `DescribePackageFilters` operations, which returns results in the next page.")
	opensearchCmd.AddCommand(opensearch_describePackagesCmd)
}
