package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getServiceEndpointCmd = &cobra.Command{
	Use:   "get-service-endpoint",
	Short: "Gets the account-specific endpoint for Configuration and Update Server (CUPS) protocol or LoRaWAN Network Server (LNS) connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getServiceEndpointCmd).Standalone()

	iotwireless_getServiceEndpointCmd.Flags().String("service-type", "", "The service type for which to get endpoint information about.")
	iotwirelessCmd.AddCommand(iotwireless_getServiceEndpointCmd)
}
