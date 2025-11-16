package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeClientVpnConnectionsCmd = &cobra.Command{
	Use:   "describe-client-vpn-connections",
	Short: "Describes active client connections and connections that have been terminated within the last 60 minutes for the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeClientVpnConnectionsCmd).Standalone()

	ec2_describeClientVpnConnectionsCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
	ec2_describeClientVpnConnectionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnConnectionsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeClientVpnConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeClientVpnConnectionsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	ec2_describeClientVpnConnectionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnConnectionsCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_describeClientVpnConnectionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeClientVpnConnectionsCmd)
}
