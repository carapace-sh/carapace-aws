package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateRouteResponseCmd = &cobra.Command{
	Use:   "update-route-response",
	Short: "Updates a RouteResponse.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateRouteResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_updateRouteResponseCmd).Standalone()

		apigatewayv2_updateRouteResponseCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_updateRouteResponseCmd.Flags().String("model-selection-expression", "", "The model selection expression for the route response.")
		apigatewayv2_updateRouteResponseCmd.Flags().String("response-models", "", "The response models for the route response.")
		apigatewayv2_updateRouteResponseCmd.Flags().String("response-parameters", "", "The route response parameters.")
		apigatewayv2_updateRouteResponseCmd.Flags().String("route-id", "", "The route ID.")
		apigatewayv2_updateRouteResponseCmd.Flags().String("route-response-id", "", "The route response ID.")
		apigatewayv2_updateRouteResponseCmd.Flags().String("route-response-key", "", "The route response key.")
		apigatewayv2_updateRouteResponseCmd.MarkFlagRequired("api-id")
		apigatewayv2_updateRouteResponseCmd.MarkFlagRequired("route-id")
		apigatewayv2_updateRouteResponseCmd.MarkFlagRequired("route-response-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateRouteResponseCmd)
}
