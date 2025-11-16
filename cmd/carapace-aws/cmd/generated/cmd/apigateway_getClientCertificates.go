package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getClientCertificatesCmd = &cobra.Command{
	Use:   "get-client-certificates",
	Short: "Gets a collection of ClientCertificate resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getClientCertificatesCmd).Standalone()

	apigateway_getClientCertificatesCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getClientCertificatesCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigatewayCmd.AddCommand(apigateway_getClientCertificatesCmd)
}
