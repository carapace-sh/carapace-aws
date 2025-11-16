package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_putIntegrationResponseCmd = &cobra.Command{
	Use:   "put-integration-response",
	Short: "Represents a put integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_putIntegrationResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_putIntegrationResponseCmd).Standalone()

		apigateway_putIntegrationResponseCmd.Flags().String("content-handling", "", "Specifies how to handle response payload content type conversions.")
		apigateway_putIntegrationResponseCmd.Flags().String("http-method", "", "Specifies a put integration response request's HTTP method.")
		apigateway_putIntegrationResponseCmd.Flags().String("resource-id", "", "Specifies a put integration response request's resource identifier.")
		apigateway_putIntegrationResponseCmd.Flags().String("response-parameters", "", "A key-value map specifying response parameters that are passed to the method response from the back end.")
		apigateway_putIntegrationResponseCmd.Flags().String("response-templates", "", "Specifies a put integration response's templates.")
		apigateway_putIntegrationResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_putIntegrationResponseCmd.Flags().String("selection-pattern", "", "Specifies the selection pattern of a put integration response.")
		apigateway_putIntegrationResponseCmd.Flags().String("status-code", "", "Specifies the status code that is used to map the integration response to an existing MethodResponse.")
		apigateway_putIntegrationResponseCmd.MarkFlagRequired("http-method")
		apigateway_putIntegrationResponseCmd.MarkFlagRequired("resource-id")
		apigateway_putIntegrationResponseCmd.MarkFlagRequired("rest-api-id")
		apigateway_putIntegrationResponseCmd.MarkFlagRequired("status-code")
	})
	apigatewayCmd.AddCommand(apigateway_putIntegrationResponseCmd)
}
