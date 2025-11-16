package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listDomainsForPackageCmd = &cobra.Command{
	Use:   "list-domains-for-package",
	Short: "Lists all Amazon OpenSearch Service domains associated with a given package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listDomainsForPackageCmd).Standalone()

	opensearch_listDomainsForPackageCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_listDomainsForPackageCmd.Flags().String("next-token", "", "If your initial `ListDomainsForPackage` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListDomainsForPackage` operations, which returns results in the next page.")
	opensearch_listDomainsForPackageCmd.Flags().String("package-id", "", "The unique identifier of the package for which to list associated domains.")
	opensearch_listDomainsForPackageCmd.MarkFlagRequired("package-id")
	opensearchCmd.AddCommand(opensearch_listDomainsForPackageCmd)
}
