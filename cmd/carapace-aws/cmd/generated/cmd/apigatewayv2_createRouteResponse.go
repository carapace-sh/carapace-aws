package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createRouteResponseCmd = &cobra.Command{
	Use:   "create-route-response",
	Short: "Creates a RouteResponse for a Route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createRouteResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_createRouteResponseCmd).Standalone()

		apigatewayv2_createRouteResponseCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_createRouteResponseCmd.Flags().String("model-selection-expression", "", "The model selection expression for the route response.")
		apigatewayv2_createRouteResponseCmd.Flags().String("response-models", "", "The response models for the route response.")
		apigatewayv2_createRouteResponseCmd.Flags().String("response-parameters", "", "The route response parameters.")
		apigatewayv2_createRouteResponseCmd.Flags().String("route-id", "", "The route ID.")
		apigatewayv2_createRouteResponseCmd.Flags().String("route-response-key", "", "The route response key.")
		apigatewayv2_createRouteResponseCmd.MarkFlagRequired("api-id")
		apigatewayv2_createRouteResponseCmd.MarkFlagRequired("route-id")
		apigatewayv2_createRouteResponseCmd.MarkFlagRequired("route-response-key")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_createRouteResponseCmd)
}
