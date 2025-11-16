package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessGatewayCertificateCmd = &cobra.Command{
	Use:   "get-wireless-gateway-certificate",
	Short: "Gets the ID of the certificate that is currently associated with a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessGatewayCertificateCmd).Standalone()

	iotwireless_getWirelessGatewayCertificateCmd.Flags().String("id", "", "The ID of the resource to get.")
	iotwireless_getWirelessGatewayCertificateCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessGatewayCertificateCmd)
}
