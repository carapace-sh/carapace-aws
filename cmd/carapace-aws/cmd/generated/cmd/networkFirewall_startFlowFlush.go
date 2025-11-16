package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_startFlowFlushCmd = &cobra.Command{
	Use:   "start-flow-flush",
	Short: "Begins the flushing of traffic from the firewall, according to the filters you define.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_startFlowFlushCmd).Standalone()

	networkFirewall_startFlowFlushCmd.Flags().String("availability-zone", "", "The ID of the Availability Zone where the firewall is located.")
	networkFirewall_startFlowFlushCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
	networkFirewall_startFlowFlushCmd.Flags().String("flow-filters", "", "Defines the scope a flow operation.")
	networkFirewall_startFlowFlushCmd.Flags().String("minimum-flow-age-in-seconds", "", "The reqested `FlowOperation` ignores flows with an age (in seconds) lower than `MinimumFlowAgeInSeconds`.")
	networkFirewall_startFlowFlushCmd.Flags().String("vpc-endpoint-association-arn", "", "The Amazon Resource Name (ARN) of a VPC endpoint association.")
	networkFirewall_startFlowFlushCmd.Flags().String("vpc-endpoint-id", "", "A unique identifier for the primary endpoint associated with a firewall.")
	networkFirewall_startFlowFlushCmd.MarkFlagRequired("firewall-arn")
	networkFirewall_startFlowFlushCmd.MarkFlagRequired("flow-filters")
	networkFirewallCmd.AddCommand(networkFirewall_startFlowFlushCmd)
}
