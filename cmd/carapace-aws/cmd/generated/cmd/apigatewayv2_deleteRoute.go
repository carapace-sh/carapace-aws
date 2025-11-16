package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteRouteCmd = &cobra.Command{
	Use:   "delete-route",
	Short: "Deletes a Route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteRouteCmd).Standalone()

	apigatewayv2_deleteRouteCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteRouteCmd.Flags().String("route-id", "", "The route ID.")
	apigatewayv2_deleteRouteCmd.MarkFlagRequired("api-id")
	apigatewayv2_deleteRouteCmd.MarkFlagRequired("route-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteRouteCmd)
}
