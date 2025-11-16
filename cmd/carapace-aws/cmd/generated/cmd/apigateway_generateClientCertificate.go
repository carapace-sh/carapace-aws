package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_generateClientCertificateCmd = &cobra.Command{
	Use:   "generate-client-certificate",
	Short: "Generates a ClientCertificate resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_generateClientCertificateCmd).Standalone()

	apigateway_generateClientCertificateCmd.Flags().String("description", "", "The description of the ClientCertificate.")
	apigateway_generateClientCertificateCmd.Flags().String("tags", "", "The key-value map of strings.")
	apigatewayCmd.AddCommand(apigateway_generateClientCertificateCmd)
}
