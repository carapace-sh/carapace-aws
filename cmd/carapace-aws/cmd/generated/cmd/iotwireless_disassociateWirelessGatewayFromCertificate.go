package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_disassociateWirelessGatewayFromCertificateCmd = &cobra.Command{
	Use:   "disassociate-wireless-gateway-from-certificate",
	Short: "Disassociates a wireless gateway from its currently associated certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_disassociateWirelessGatewayFromCertificateCmd).Standalone()

	iotwireless_disassociateWirelessGatewayFromCertificateCmd.Flags().String("id", "", "The ID of the resource to update.")
	iotwireless_disassociateWirelessGatewayFromCertificateCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_disassociateWirelessGatewayFromCertificateCmd)
}
