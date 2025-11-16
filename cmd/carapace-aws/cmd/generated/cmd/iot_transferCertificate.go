package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_transferCertificateCmd = &cobra.Command{
	Use:   "transfer-certificate",
	Short: "Transfers the specified certificate to the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_transferCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_transferCertificateCmd).Standalone()

		iot_transferCertificateCmd.Flags().String("certificate-id", "", "The ID of the certificate.")
		iot_transferCertificateCmd.Flags().String("target-aws-account", "", "The Amazon Web Services account.")
		iot_transferCertificateCmd.Flags().String("transfer-message", "", "The transfer message.")
		iot_transferCertificateCmd.MarkFlagRequired("certificate-id")
		iot_transferCertificateCmd.MarkFlagRequired("target-aws-account")
	})
	iotCmd.AddCommand(iot_transferCertificateCmd)
}
