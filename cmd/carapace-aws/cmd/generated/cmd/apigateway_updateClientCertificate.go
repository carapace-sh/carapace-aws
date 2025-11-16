package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateClientCertificateCmd = &cobra.Command{
	Use:   "update-client-certificate",
	Short: "Changes information about an ClientCertificate resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateClientCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateClientCertificateCmd).Standalone()

		apigateway_updateClientCertificateCmd.Flags().String("client-certificate-id", "", "The identifier of the ClientCertificate resource to be updated.")
		apigateway_updateClientCertificateCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateClientCertificateCmd.MarkFlagRequired("client-certificate-id")
	})
	apigatewayCmd.AddCommand(apigateway_updateClientCertificateCmd)
}
