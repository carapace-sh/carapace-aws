package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_associateWirelessGatewayWithCertificateCmd = &cobra.Command{
	Use:   "associate-wireless-gateway-with-certificate",
	Short: "Associates a wireless gateway with a certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_associateWirelessGatewayWithCertificateCmd).Standalone()

	iotwireless_associateWirelessGatewayWithCertificateCmd.Flags().String("id", "", "The ID of the resource to update.")
	iotwireless_associateWirelessGatewayWithCertificateCmd.Flags().String("iot-certificate-id", "", "The ID of the certificate to associate with the wireless gateway.")
	iotwireless_associateWirelessGatewayWithCertificateCmd.MarkFlagRequired("id")
	iotwireless_associateWirelessGatewayWithCertificateCmd.MarkFlagRequired("iot-certificate-id")
	iotwirelessCmd.AddCommand(iotwireless_associateWirelessGatewayWithCertificateCmd)
}
