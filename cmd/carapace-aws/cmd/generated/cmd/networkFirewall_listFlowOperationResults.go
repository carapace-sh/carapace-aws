package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listFlowOperationResultsCmd = &cobra.Command{
	Use:   "list-flow-operation-results",
	Short: "Returns the results of a specific flow operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listFlowOperationResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_listFlowOperationResultsCmd).Standalone()

		networkFirewall_listFlowOperationResultsCmd.Flags().String("availability-zone", "", "The ID of the Availability Zone where the firewall is located.")
		networkFirewall_listFlowOperationResultsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_listFlowOperationResultsCmd.Flags().String("flow-operation-id", "", "A unique identifier for the flow operation.")
		networkFirewall_listFlowOperationResultsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
		networkFirewall_listFlowOperationResultsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
		networkFirewall_listFlowOperationResultsCmd.Flags().String("vpc-endpoint-association-arn", "", "The Amazon Resource Name (ARN) of a VPC endpoint association.")
		networkFirewall_listFlowOperationResultsCmd.Flags().String("vpc-endpoint-id", "", "A unique identifier for the primary endpoint associated with a firewall.")
		networkFirewall_listFlowOperationResultsCmd.MarkFlagRequired("firewall-arn")
		networkFirewall_listFlowOperationResultsCmd.MarkFlagRequired("flow-operation-id")
	})
	networkFirewallCmd.AddCommand(networkFirewall_listFlowOperationResultsCmd)
}
