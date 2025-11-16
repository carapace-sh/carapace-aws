package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_startFlowCaptureCmd = &cobra.Command{
	Use:   "start-flow-capture",
	Short: "Begins capturing the flows in a firewall, according to the filters you define.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_startFlowCaptureCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_startFlowCaptureCmd).Standalone()

		networkFirewall_startFlowCaptureCmd.Flags().String("availability-zone", "", "The ID of the Availability Zone where the firewall is located.")
		networkFirewall_startFlowCaptureCmd.Flags().String("firewall-arn", "", "The Amazon Resource Name (ARN) of the firewall.")
		networkFirewall_startFlowCaptureCmd.Flags().String("flow-filters", "", "Defines the scope a flow operation.")
		networkFirewall_startFlowCaptureCmd.Flags().String("minimum-flow-age-in-seconds", "", "The reqested `FlowOperation` ignores flows with an age (in seconds) lower than `MinimumFlowAgeInSeconds`.")
		networkFirewall_startFlowCaptureCmd.Flags().String("vpc-endpoint-association-arn", "", "The Amazon Resource Name (ARN) of a VPC endpoint association.")
		networkFirewall_startFlowCaptureCmd.Flags().String("vpc-endpoint-id", "", "A unique identifier for the primary endpoint associated with a firewall.")
		networkFirewall_startFlowCaptureCmd.MarkFlagRequired("firewall-arn")
		networkFirewall_startFlowCaptureCmd.MarkFlagRequired("flow-filters")
	})
	networkFirewallCmd.AddCommand(networkFirewall_startFlowCaptureCmd)
}
