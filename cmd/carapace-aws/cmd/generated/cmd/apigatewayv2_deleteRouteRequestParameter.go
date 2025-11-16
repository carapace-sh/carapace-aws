package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteRouteRequestParameterCmd = &cobra.Command{
	Use:   "delete-route-request-parameter",
	Short: "Deletes a route request parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteRouteRequestParameterCmd).Standalone()

	apigatewayv2_deleteRouteRequestParameterCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteRouteRequestParameterCmd.Flags().String("request-parameter-key", "", "The route request parameter key.")
	apigatewayv2_deleteRouteRequestParameterCmd.Flags().String("route-id", "", "The route ID.")
	apigatewayv2_deleteRouteRequestParameterCmd.MarkFlagRequired("api-id")
	apigatewayv2_deleteRouteRequestParameterCmd.MarkFlagRequired("request-parameter-key")
	apigatewayv2_deleteRouteRequestParameterCmd.MarkFlagRequired("route-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteRouteRequestParameterCmd)
}
