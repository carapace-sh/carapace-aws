package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_rejectCertificateTransferCmd = &cobra.Command{
	Use:   "reject-certificate-transfer",
	Short: "Rejects a pending certificate transfer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_rejectCertificateTransferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_rejectCertificateTransferCmd).Standalone()

		iot_rejectCertificateTransferCmd.Flags().String("certificate-id", "", "The ID of the certificate.")
		iot_rejectCertificateTransferCmd.Flags().String("reject-reason", "", "The reason the certificate transfer was rejected.")
		iot_rejectCertificateTransferCmd.MarkFlagRequired("certificate-id")
	})
	iotCmd.AddCommand(iot_rejectCertificateTransferCmd)
}
