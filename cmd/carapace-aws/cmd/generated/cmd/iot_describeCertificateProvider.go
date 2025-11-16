package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeCertificateProviderCmd = &cobra.Command{
	Use:   "describe-certificate-provider",
	Short: "Describes a certificate provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeCertificateProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeCertificateProviderCmd).Standalone()

		iot_describeCertificateProviderCmd.Flags().String("certificate-provider-name", "", "The name of the certificate provider.")
		iot_describeCertificateProviderCmd.MarkFlagRequired("certificate-provider-name")
	})
	iotCmd.AddCommand(iot_describeCertificateProviderCmd)
}
