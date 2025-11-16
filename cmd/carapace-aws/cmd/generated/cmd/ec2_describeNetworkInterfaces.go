package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkInterfacesCmd = &cobra.Command{
	Use:   "describe-network-interfaces",
	Short: "Describes the specified network interfaces or all your network interfaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkInterfacesCmd).Standalone()

	ec2_describeNetworkInterfacesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkInterfacesCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeNetworkInterfacesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeNetworkInterfacesCmd.Flags().String("network-interface-ids", "", "The network interface IDs.")
	ec2_describeNetworkInterfacesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeNetworkInterfacesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkInterfacesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeNetworkInterfacesCmd)
}
