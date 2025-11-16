package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getVpnConnectionDeviceTypesCmd = &cobra.Command{
	Use:   "get-vpn-connection-device-types",
	Short: "Obtain a list of customer gateway devices for which sample configuration files can be provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getVpnConnectionDeviceTypesCmd).Standalone()

	ec2_getVpnConnectionDeviceTypesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getVpnConnectionDeviceTypesCmd.Flags().String("max-results", "", "The maximum number of results returned by `GetVpnConnectionDeviceTypes` in paginated output.")
	ec2_getVpnConnectionDeviceTypesCmd.Flags().String("next-token", "", "The `NextToken` value returned from a previous paginated `GetVpnConnectionDeviceTypes` request where `MaxResults` was used and the results exceeded the value of that parameter.")
	ec2_getVpnConnectionDeviceTypesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getVpnConnectionDeviceTypesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getVpnConnectionDeviceTypesCmd)
}
