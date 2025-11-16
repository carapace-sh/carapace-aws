package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getRoutesCmd = &cobra.Command{
	Use:   "get-routes",
	Short: "Gets the Routes for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getRoutesCmd).Standalone()

		apigatewayv2_getRoutesCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getRoutesCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getRoutesCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
		apigatewayv2_getRoutesCmd.MarkFlagRequired("api-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getRoutesCmd)
}
