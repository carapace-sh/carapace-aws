package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteAuthorizerCmd = &cobra.Command{
	Use:   "delete-authorizer",
	Short: "Deletes an existing Authorizer resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteAuthorizerCmd).Standalone()

	apigateway_deleteAuthorizerCmd.Flags().String("authorizer-id", "", "The identifier of the Authorizer resource.")
	apigateway_deleteAuthorizerCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_deleteAuthorizerCmd.MarkFlagRequired("authorizer-id")
	apigateway_deleteAuthorizerCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_deleteAuthorizerCmd)
}
