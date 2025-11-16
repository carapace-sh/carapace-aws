package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getClientCertificateCmd = &cobra.Command{
	Use:   "get-client-certificate",
	Short: "Gets information about the current ClientCertificate resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getClientCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getClientCertificateCmd).Standalone()

		apigateway_getClientCertificateCmd.Flags().String("client-certificate-id", "", "The identifier of the ClientCertificate resource to be described.")
		apigateway_getClientCertificateCmd.MarkFlagRequired("client-certificate-id")
	})
	apigatewayCmd.AddCommand(apigateway_getClientCertificateCmd)
}
