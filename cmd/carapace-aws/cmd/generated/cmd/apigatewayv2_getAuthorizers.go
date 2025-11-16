package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getAuthorizersCmd = &cobra.Command{
	Use:   "get-authorizers",
	Short: "Gets the Authorizers for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getAuthorizersCmd).Standalone()

	apigatewayv2_getAuthorizersCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getAuthorizersCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
	apigatewayv2_getAuthorizersCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
	apigatewayv2_getAuthorizersCmd.MarkFlagRequired("api-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getAuthorizersCmd)
}
