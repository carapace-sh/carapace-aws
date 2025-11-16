package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_listDataIntegrationFlowExecutionsCmd = &cobra.Command{
	Use:   "list-data-integration-flow-executions",
	Short: "List flow executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_listDataIntegrationFlowExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_listDataIntegrationFlowExecutionsCmd).Standalone()

		supplychain_listDataIntegrationFlowExecutionsCmd.Flags().String("flow-name", "", "The flow name.")
		supplychain_listDataIntegrationFlowExecutionsCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
		supplychain_listDataIntegrationFlowExecutionsCmd.Flags().String("max-results", "", "The number to specify the max number of flow executions to fetch in this paginated request.")
		supplychain_listDataIntegrationFlowExecutionsCmd.Flags().String("next-token", "", "The pagination token to fetch next page of flow executions.")
		supplychain_listDataIntegrationFlowExecutionsCmd.MarkFlagRequired("flow-name")
		supplychain_listDataIntegrationFlowExecutionsCmd.MarkFlagRequired("instance-id")
	})
	supplychainCmd.AddCommand(supplychain_listDataIntegrationFlowExecutionsCmd)
}
