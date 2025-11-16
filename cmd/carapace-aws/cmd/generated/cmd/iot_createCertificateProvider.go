package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createCertificateProviderCmd = &cobra.Command{
	Use:   "create-certificate-provider",
	Short: "Creates an Amazon Web Services IoT Core certificate provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createCertificateProviderCmd).Standalone()

	iot_createCertificateProviderCmd.Flags().String("account-default-for-operations", "", "A list of the operations that the certificate provider will use to generate certificates.")
	iot_createCertificateProviderCmd.Flags().String("certificate-provider-name", "", "The name of the certificate provider.")
	iot_createCertificateProviderCmd.Flags().String("client-token", "", "A string that you can optionally pass in the `CreateCertificateProvider` request to make sure the request is idempotent.")
	iot_createCertificateProviderCmd.Flags().String("lambda-function-arn", "", "The ARN of the Lambda function that defines the authentication logic.")
	iot_createCertificateProviderCmd.Flags().String("tags", "", "Metadata which can be used to manage the certificate provider.")
	iot_createCertificateProviderCmd.MarkFlagRequired("account-default-for-operations")
	iot_createCertificateProviderCmd.MarkFlagRequired("certificate-provider-name")
	iot_createCertificateProviderCmd.MarkFlagRequired("lambda-function-arn")
	iotCmd.AddCommand(iot_createCertificateProviderCmd)
}
