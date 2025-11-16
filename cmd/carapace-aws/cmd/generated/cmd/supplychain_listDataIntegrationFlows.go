package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_listDataIntegrationFlowsCmd = &cobra.Command{
	Use:   "list-data-integration-flows",
	Short: "Enables you to programmatically list all data pipelines for the provided Amazon Web Services Supply Chain instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_listDataIntegrationFlowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_listDataIntegrationFlowsCmd).Standalone()

		supplychain_listDataIntegrationFlowsCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
		supplychain_listDataIntegrationFlowsCmd.Flags().String("max-results", "", "Specify the maximum number of DataIntegrationFlows to fetch in one paginated request.")
		supplychain_listDataIntegrationFlowsCmd.Flags().String("next-token", "", "The pagination token to fetch the next page of the DataIntegrationFlows.")
		supplychain_listDataIntegrationFlowsCmd.MarkFlagRequired("instance-id")
	})
	supplychainCmd.AddCommand(supplychain_listDataIntegrationFlowsCmd)
}
