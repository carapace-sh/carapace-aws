package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateCertificateProviderCmd = &cobra.Command{
	Use:   "update-certificate-provider",
	Short: "Updates a certificate provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateCertificateProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateCertificateProviderCmd).Standalone()

		iot_updateCertificateProviderCmd.Flags().String("account-default-for-operations", "", "A list of the operations that the certificate provider will use to generate certificates.")
		iot_updateCertificateProviderCmd.Flags().String("certificate-provider-name", "", "The name of the certificate provider.")
		iot_updateCertificateProviderCmd.Flags().String("lambda-function-arn", "", "The Lambda function ARN that's associated with the certificate provider.")
		iot_updateCertificateProviderCmd.MarkFlagRequired("certificate-provider-name")
	})
	iotCmd.AddCommand(iot_updateCertificateProviderCmd)
}
