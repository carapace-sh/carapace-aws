package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listPackagesForDomainCmd = &cobra.Command{
	Use:   "list-packages-for-domain",
	Short: "Lists all packages associated with an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listPackagesForDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_listPackagesForDomainCmd).Standalone()

		opensearch_listPackagesForDomainCmd.Flags().String("domain-name", "", "The name of the domain for which you want to list associated packages.")
		opensearch_listPackagesForDomainCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearch_listPackagesForDomainCmd.Flags().String("next-token", "", "If your initial `ListPackagesForDomain` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListPackagesForDomain` operations, which returns results in the next page.")
		opensearch_listPackagesForDomainCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_listPackagesForDomainCmd)
}
