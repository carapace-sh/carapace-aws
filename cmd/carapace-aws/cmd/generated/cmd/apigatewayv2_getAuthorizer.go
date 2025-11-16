package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getAuthorizerCmd = &cobra.Command{
	Use:   "get-authorizer",
	Short: "Gets an Authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getAuthorizerCmd).Standalone()

		apigatewayv2_getAuthorizerCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getAuthorizerCmd.Flags().String("authorizer-id", "", "The authorizer identifier.")
		apigatewayv2_getAuthorizerCmd.MarkFlagRequired("api-id")
		apigatewayv2_getAuthorizerCmd.MarkFlagRequired("authorizer-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getAuthorizerCmd)
}
