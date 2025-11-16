package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteClientCertificateCmd = &cobra.Command{
	Use:   "delete-client-certificate",
	Short: "Deletes the ClientCertificate resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteClientCertificateCmd).Standalone()

	apigateway_deleteClientCertificateCmd.Flags().String("client-certificate-id", "", "The identifier of the ClientCertificate resource to be deleted.")
	apigateway_deleteClientCertificateCmd.MarkFlagRequired("client-certificate-id")
	apigatewayCmd.AddCommand(apigateway_deleteClientCertificateCmd)
}
