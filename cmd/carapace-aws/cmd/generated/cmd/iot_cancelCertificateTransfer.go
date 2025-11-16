package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_cancelCertificateTransferCmd = &cobra.Command{
	Use:   "cancel-certificate-transfer",
	Short: "Cancels a pending transfer for the specified certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_cancelCertificateTransferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_cancelCertificateTransferCmd).Standalone()

		iot_cancelCertificateTransferCmd.Flags().String("certificate-id", "", "The ID of the certificate.")
		iot_cancelCertificateTransferCmd.MarkFlagRequired("certificate-id")
	})
	iotCmd.AddCommand(iot_cancelCertificateTransferCmd)
}
