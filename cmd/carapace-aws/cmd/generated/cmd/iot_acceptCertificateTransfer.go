package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_acceptCertificateTransferCmd = &cobra.Command{
	Use:   "accept-certificate-transfer",
	Short: "Accepts a pending certificate transfer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_acceptCertificateTransferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_acceptCertificateTransferCmd).Standalone()

		iot_acceptCertificateTransferCmd.Flags().String("certificate-id", "", "The ID of the certificate.")
		iot_acceptCertificateTransferCmd.Flags().String("set-as-active", "", "Specifies whether the certificate is active.")
		iot_acceptCertificateTransferCmd.MarkFlagRequired("certificate-id")
	})
	iotCmd.AddCommand(iot_acceptCertificateTransferCmd)
}
