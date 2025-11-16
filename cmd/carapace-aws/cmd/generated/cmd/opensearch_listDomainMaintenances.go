package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listDomainMaintenancesCmd = &cobra.Command{
	Use:   "list-domain-maintenances",
	Short: "A list of maintenance actions for the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listDomainMaintenancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_listDomainMaintenancesCmd).Standalone()

		opensearch_listDomainMaintenancesCmd.Flags().String("action", "", "The name of the action.")
		opensearch_listDomainMaintenancesCmd.Flags().String("domain-name", "", "The name of the domain.")
		opensearch_listDomainMaintenancesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearch_listDomainMaintenancesCmd.Flags().String("next-token", "", "If your initial `ListDomainMaintenances` operation returns a `nextToken`, include the returned `nextToken` in subsequent `ListDomainMaintenances` operations, which returns results in the next page.")
		opensearch_listDomainMaintenancesCmd.Flags().String("status", "", "The status of the action.")
		opensearch_listDomainMaintenancesCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_listDomainMaintenancesCmd)
}
