package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeCertificateCmd = &cobra.Command{
	Use:   "describe-certificate",
	Short: "Gets information about the specified certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeCertificateCmd).Standalone()

		iot_describeCertificateCmd.Flags().String("certificate-id", "", "The ID of the certificate.")
		iot_describeCertificateCmd.MarkFlagRequired("certificate-id")
	})
	iotCmd.AddCommand(iot_describeCertificateCmd)
}
