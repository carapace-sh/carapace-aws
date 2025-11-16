package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateRouteCmd = &cobra.Command{
	Use:   "update-route",
	Short: "Updates a Route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_updateRouteCmd).Standalone()

		apigatewayv2_updateRouteCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_updateRouteCmd.Flags().String("api-key-required", "", "Specifies whether an API key is required for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("authorization-scopes", "", "The authorization scopes supported by this route.")
		apigatewayv2_updateRouteCmd.Flags().String("authorization-type", "", "The authorization type for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("authorizer-id", "", "The identifier of the Authorizer resource to be associated with this route.")
		apigatewayv2_updateRouteCmd.Flags().String("model-selection-expression", "", "The model selection expression for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("operation-name", "", "The operation name for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("request-models", "", "The request models for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("request-parameters", "", "The request parameters for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("route-id", "", "The route ID.")
		apigatewayv2_updateRouteCmd.Flags().String("route-key", "", "The route key for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("route-response-selection-expression", "", "The route response selection expression for the route.")
		apigatewayv2_updateRouteCmd.Flags().String("target", "", "The target for the route.")
		apigatewayv2_updateRouteCmd.MarkFlagRequired("api-id")
		apigatewayv2_updateRouteCmd.MarkFlagRequired("route-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateRouteCmd)
}
