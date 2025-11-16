package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listFlowOperationsCmd = &cobra.Command{
	Use:   "list-flow-operations",
	Short: "Returns a list of all flow operations ran in a specific firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listFlowOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_listFlowOperationsCmd).Standalone()

		networkFirewall_listFlowOperationsCmd.Flags().String("availability-zone", "", "The ID of the Availability Zone where the firewall is located.")
		networkFirewall_listFlowOperationsCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_listFlowOperationsCmd.Flags().String("flow-operation-type", "", "An optional string that defines whether any or all operation types are returned.")
		networkFirewall_listFlowOperationsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
		networkFirewall_listFlowOperationsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
		networkFirewall_listFlowOperationsCmd.Flags().String("vpc-endpoint-association-arn", "", "The Amazon Resource Name (ARN) of a VPC endpoint association.")
		networkFirewall_listFlowOperationsCmd.Flags().String("vpc-endpoint-id", "", "A unique identifier for the primary endpoint associated with a firewall.")
		networkFirewall_listFlowOperationsCmd.MarkFlagRequired("firewall-arn")
	})
	networkFirewallCmd.AddCommand(networkFirewall_listFlowOperationsCmd)
}
