package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLocalGatewayVirtualInterfacesCmd = &cobra.Command{
	Use:   "describe-local-gateway-virtual-interfaces",
	Short: "Describes the specified local gateway virtual interfaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLocalGatewayVirtualInterfacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeLocalGatewayVirtualInterfacesCmd).Standalone()

		ec2_describeLocalGatewayVirtualInterfacesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayVirtualInterfacesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeLocalGatewayVirtualInterfacesCmd.Flags().String("local-gateway-virtual-interface-ids", "", "The IDs of the virtual interfaces.")
		ec2_describeLocalGatewayVirtualInterfacesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeLocalGatewayVirtualInterfacesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeLocalGatewayVirtualInterfacesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeLocalGatewayVirtualInterfacesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeLocalGatewayVirtualInterfacesCmd)
}
