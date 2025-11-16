package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeFlowOperationCmd = &cobra.Command{
	Use:   "describe-flow-operation",
	Short: "Returns key information about a specific flow operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeFlowOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_describeFlowOperationCmd).Standalone()

		networkFirewall_describeFlowOperationCmd.Flags().String("availability-zone", "", "The ID of the Availability Zone where the firewall is located.")
		networkFirewall_describeFlowOperationCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_describeFlowOperationCmd.Flags().String("flow-operation-id", "", "A unique identifier for the flow operation.")
		networkFirewall_describeFlowOperationCmd.Flags().String("vpc-endpoint-association-arn", "", "The Amazon Resource Name (ARN) of a VPC endpoint association.")
		networkFirewall_describeFlowOperationCmd.Flags().String("vpc-endpoint-id", "", "A unique identifier for the primary endpoint associated with a firewall.")
		networkFirewall_describeFlowOperationCmd.MarkFlagRequired("firewall-arn")
		networkFirewall_describeFlowOperationCmd.MarkFlagRequired("flow-operation-id")
	})
	networkFirewallCmd.AddCommand(networkFirewall_describeFlowOperationCmd)
}
