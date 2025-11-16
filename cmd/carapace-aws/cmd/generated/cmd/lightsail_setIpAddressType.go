package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_setIpAddressTypeCmd = &cobra.Command{
	Use:   "set-ip-address-type",
	Short: "Sets the IP address type for an Amazon Lightsail resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_setIpAddressTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_setIpAddressTypeCmd).Standalone()

		lightsail_setIpAddressTypeCmd.Flags().String("accept-bundle-update", "", "Required parameter to accept the instance bundle update when changing to, and from, IPv6-only.")
		lightsail_setIpAddressTypeCmd.Flags().String("ip-address-type", "", "The IP address type to set for the specified resource.")
		lightsail_setIpAddressTypeCmd.Flags().String("resource-name", "", "The name of the resource for which to set the IP address type.")
		lightsail_setIpAddressTypeCmd.Flags().String("resource-type", "", "The resource type.")
		lightsail_setIpAddressTypeCmd.MarkFlagRequired("ip-address-type")
		lightsail_setIpAddressTypeCmd.MarkFlagRequired("resource-name")
		lightsail_setIpAddressTypeCmd.MarkFlagRequired("resource-type")
	})
	lightsailCmd.AddCommand(lightsail_setIpAddressTypeCmd)
}
