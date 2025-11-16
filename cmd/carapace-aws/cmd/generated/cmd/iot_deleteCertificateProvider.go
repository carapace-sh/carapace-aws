package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteCertificateProviderCmd = &cobra.Command{
	Use:   "delete-certificate-provider",
	Short: "Deletes a certificate provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteCertificateProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteCertificateProviderCmd).Standalone()

		iot_deleteCertificateProviderCmd.Flags().String("certificate-provider-name", "", "The name of the certificate provider.")
		iot_deleteCertificateProviderCmd.MarkFlagRequired("certificate-provider-name")
	})
	iotCmd.AddCommand(iot_deleteCertificateProviderCmd)
}
