package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_exportClientVpnClientConfigurationCmd = &cobra.Command{
	Use:   "export-client-vpn-client-configuration",
	Short: "Downloads the contents of the Client VPN endpoint configuration file for the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_exportClientVpnClientConfigurationCmd).Standalone()

	ec2_exportClientVpnClientConfigurationCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
	ec2_exportClientVpnClientConfigurationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportClientVpnClientConfigurationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportClientVpnClientConfigurationCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_exportClientVpnClientConfigurationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_exportClientVpnClientConfigurationCmd)
}
