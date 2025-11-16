package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeClientVpnEndpointsCmd = &cobra.Command{
	Use:   "describe-client-vpn-endpoints",
	Short: "Describes one or more Client VPN endpoints in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeClientVpnEndpointsCmd).Standalone()

	ec2_describeClientVpnEndpointsCmd.Flags().String("client-vpn-endpoint-ids", "", "The ID of the Client VPN endpoint.")
	ec2_describeClientVpnEndpointsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnEndpointsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeClientVpnEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeClientVpnEndpointsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	ec2_describeClientVpnEndpointsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnEndpointsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeClientVpnEndpointsCmd)
}
