package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeClientVpnRoutesCmd = &cobra.Command{
	Use:   "describe-client-vpn-routes",
	Short: "Describes the routes for the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeClientVpnRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeClientVpnRoutesCmd).Standalone()

		ec2_describeClientVpnRoutesCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
		ec2_describeClientVpnRoutesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeClientVpnRoutesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeClientVpnRoutesCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
		ec2_describeClientVpnRoutesCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
		ec2_describeClientVpnRoutesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeClientVpnRoutesCmd.MarkFlagRequired("client-vpn-endpoint-id")
		ec2_describeClientVpnRoutesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeClientVpnRoutesCmd)
}
