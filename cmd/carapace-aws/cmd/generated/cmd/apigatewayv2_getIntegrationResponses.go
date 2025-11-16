package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getIntegrationResponsesCmd = &cobra.Command{
	Use:   "get-integration-responses",
	Short: "Gets the IntegrationResponses for an Integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getIntegrationResponsesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getIntegrationResponsesCmd).Standalone()

		apigatewayv2_getIntegrationResponsesCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getIntegrationResponsesCmd.Flags().String("integration-id", "", "The integration ID.")
		apigatewayv2_getIntegrationResponsesCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getIntegrationResponsesCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
		apigatewayv2_getIntegrationResponsesCmd.MarkFlagRequired("api-id")
		apigatewayv2_getIntegrationResponsesCmd.MarkFlagRequired("integration-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getIntegrationResponsesCmd)
}
