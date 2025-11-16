package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getIntegrationResponseCmd = &cobra.Command{
	Use:   "get-integration-response",
	Short: "Represents a get integration response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getIntegrationResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getIntegrationResponseCmd).Standalone()

		apigateway_getIntegrationResponseCmd.Flags().String("http-method", "", "Specifies a get integration response request's HTTP method.")
		apigateway_getIntegrationResponseCmd.Flags().String("resource-id", "", "Specifies a get integration response request's resource identifier.")
		apigateway_getIntegrationResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getIntegrationResponseCmd.Flags().String("status-code", "", "Specifies a get integration response request's status code.")
		apigateway_getIntegrationResponseCmd.MarkFlagRequired("http-method")
		apigateway_getIntegrationResponseCmd.MarkFlagRequired("resource-id")
		apigateway_getIntegrationResponseCmd.MarkFlagRequired("rest-api-id")
		apigateway_getIntegrationResponseCmd.MarkFlagRequired("status-code")
	})
	apigatewayCmd.AddCommand(apigateway_getIntegrationResponseCmd)
}
