package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateAuthorizerCmd = &cobra.Command{
	Use:   "update-authorizer",
	Short: "Updates an existing Authorizer resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateAuthorizerCmd).Standalone()

		apigateway_updateAuthorizerCmd.Flags().String("authorizer-id", "", "The identifier of the Authorizer resource.")
		apigateway_updateAuthorizerCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateAuthorizerCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_updateAuthorizerCmd.MarkFlagRequired("authorizer-id")
		apigateway_updateAuthorizerCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_updateAuthorizerCmd)
}
