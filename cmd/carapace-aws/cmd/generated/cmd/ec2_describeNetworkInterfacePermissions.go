package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkInterfacePermissionsCmd = &cobra.Command{
	Use:   "describe-network-interface-permissions",
	Short: "Describes the permissions for your network interfaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkInterfacePermissionsCmd).Standalone()

	ec2_describeNetworkInterfacePermissionsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeNetworkInterfacePermissionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeNetworkInterfacePermissionsCmd.Flags().String("network-interface-permission-ids", "", "The network interface permission IDs.")
	ec2_describeNetworkInterfacePermissionsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2Cmd.AddCommand(ec2_describeNetworkInterfacePermissionsCmd)
}
