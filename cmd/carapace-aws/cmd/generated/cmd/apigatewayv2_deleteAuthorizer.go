package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteAuthorizerCmd = &cobra.Command{
	Use:   "delete-authorizer",
	Short: "Deletes an Authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteAuthorizerCmd).Standalone()

	apigatewayv2_deleteAuthorizerCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteAuthorizerCmd.Flags().String("authorizer-id", "", "The authorizer identifier.")
	apigatewayv2_deleteAuthorizerCmd.MarkFlagRequired("api-id")
	apigatewayv2_deleteAuthorizerCmd.MarkFlagRequired("authorizer-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteAuthorizerCmd)
}
