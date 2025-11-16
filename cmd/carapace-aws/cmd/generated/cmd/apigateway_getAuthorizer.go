package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getAuthorizerCmd = &cobra.Command{
	Use:   "get-authorizer",
	Short: "Describe an existing Authorizer resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getAuthorizerCmd).Standalone()

	apigateway_getAuthorizerCmd.Flags().String("authorizer-id", "", "The identifier of the Authorizer resource.")
	apigateway_getAuthorizerCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getAuthorizerCmd.MarkFlagRequired("authorizer-id")
	apigateway_getAuthorizerCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getAuthorizerCmd)
}
