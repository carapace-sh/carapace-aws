package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getIntegrationCmd = &cobra.Command{
	Use:   "get-integration",
	Short: "Get the integration settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getIntegrationCmd).Standalone()

		apigateway_getIntegrationCmd.Flags().String("http-method", "", "Specifies a get integration request's HTTP method.")
		apigateway_getIntegrationCmd.Flags().String("resource-id", "", "Specifies a get integration request's resource identifier")
		apigateway_getIntegrationCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getIntegrationCmd.MarkFlagRequired("http-method")
		apigateway_getIntegrationCmd.MarkFlagRequired("resource-id")
		apigateway_getIntegrationCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getIntegrationCmd)
}
