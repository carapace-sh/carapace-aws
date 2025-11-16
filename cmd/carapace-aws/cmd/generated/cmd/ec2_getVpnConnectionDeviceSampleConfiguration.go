package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getVpnConnectionDeviceSampleConfigurationCmd = &cobra.Command{
	Use:   "get-vpn-connection-device-sample-configuration",
	Short: "Download an Amazon Web Services-provided sample configuration file to be used with the customer gateway device specified for your Site-to-Site VPN connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getVpnConnectionDeviceSampleConfigurationCmd).Standalone()

	ec2_getVpnConnectionDeviceSampleConfigurationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.Flags().String("internet-key-exchange-version", "", "The IKE version to be used in the sample configuration file for your customer gateway device.")
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.Flags().String("sample-type", "", "The type of sample configuration to generate.")
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.Flags().String("vpn-connection-device-type-id", "", "Device identifier provided by the `GetVpnConnectionDeviceTypes` API.")
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.Flags().String("vpn-connection-id", "", "The `VpnConnectionId` specifies the Site-to-Site VPN connection used for the sample configuration.")
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.Flag("no-dry-run").Hidden = true
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.MarkFlagRequired("vpn-connection-device-type-id")
	ec2_getVpnConnectionDeviceSampleConfigurationCmd.MarkFlagRequired("vpn-connection-id")
	ec2Cmd.AddCommand(ec2_getVpnConnectionDeviceSampleConfigurationCmd)
}
