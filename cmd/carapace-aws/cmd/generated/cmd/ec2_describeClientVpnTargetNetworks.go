package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeClientVpnTargetNetworksCmd = &cobra.Command{
	Use:   "describe-client-vpn-target-networks",
	Short: "Describes the target networks associated with the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeClientVpnTargetNetworksCmd).Standalone()

	ec2_describeClientVpnTargetNetworksCmd.Flags().String("association-ids", "", "The IDs of the target network associations.")
	ec2_describeClientVpnTargetNetworksCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
	ec2_describeClientVpnTargetNetworksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnTargetNetworksCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeClientVpnTargetNetworksCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeClientVpnTargetNetworksCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	ec2_describeClientVpnTargetNetworksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnTargetNetworksCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_describeClientVpnTargetNetworksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeClientVpnTargetNetworksCmd)
}
