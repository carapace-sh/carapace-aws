package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteIntegrationResponseCmd = &cobra.Command{
	Use:   "delete-integration-response",
	Short: "Represents a delete integration response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteIntegrationResponseCmd).Standalone()

	apigateway_deleteIntegrationResponseCmd.Flags().String("http-method", "", "Specifies a delete integration response request's HTTP method.")
	apigateway_deleteIntegrationResponseCmd.Flags().String("resource-id", "", "Specifies a delete integration response request's resource identifier.")
	apigateway_deleteIntegrationResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_deleteIntegrationResponseCmd.Flags().String("status-code", "", "Specifies a delete integration response request's status code.")
	apigateway_deleteIntegrationResponseCmd.MarkFlagRequired("http-method")
	apigateway_deleteIntegrationResponseCmd.MarkFlagRequired("resource-id")
	apigateway_deleteIntegrationResponseCmd.MarkFlagRequired("rest-api-id")
	apigateway_deleteIntegrationResponseCmd.MarkFlagRequired("status-code")
	apigatewayCmd.AddCommand(apigateway_deleteIntegrationResponseCmd)
}
