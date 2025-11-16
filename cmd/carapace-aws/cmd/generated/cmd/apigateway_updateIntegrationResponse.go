package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateIntegrationResponseCmd = &cobra.Command{
	Use:   "update-integration-response",
	Short: "Represents an update integration response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateIntegrationResponseCmd).Standalone()

	apigateway_updateIntegrationResponseCmd.Flags().String("http-method", "", "Specifies an update integration response request's HTTP method.")
	apigateway_updateIntegrationResponseCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateIntegrationResponseCmd.Flags().String("resource-id", "", "Specifies an update integration response request's resource identifier.")
	apigateway_updateIntegrationResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateIntegrationResponseCmd.Flags().String("status-code", "", "Specifies an update integration response request's status code.")
	apigateway_updateIntegrationResponseCmd.MarkFlagRequired("http-method")
	apigateway_updateIntegrationResponseCmd.MarkFlagRequired("resource-id")
	apigateway_updateIntegrationResponseCmd.MarkFlagRequired("rest-api-id")
	apigateway_updateIntegrationResponseCmd.MarkFlagRequired("status-code")
	apigatewayCmd.AddCommand(apigateway_updateIntegrationResponseCmd)
}
