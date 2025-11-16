package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createAuthorizerCmd = &cobra.Command{
	Use:   "create-authorizer",
	Short: "Adds a new Authorizer resource to an existing RestApi resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createAuthorizerCmd).Standalone()

		apigateway_createAuthorizerCmd.Flags().String("auth-type", "", "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.")
		apigateway_createAuthorizerCmd.Flags().String("authorizer-credentials", "", "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.")
		apigateway_createAuthorizerCmd.Flags().String("authorizer-result-ttl-in-seconds", "", "The TTL in seconds of cached authorizer results.")
		apigateway_createAuthorizerCmd.Flags().String("authorizer-uri", "", "Specifies the authorizer's Uniform Resource Identifier (URI).")
		apigateway_createAuthorizerCmd.Flags().String("identity-source", "", "The identity source for which authorization is requested.")
		apigateway_createAuthorizerCmd.Flags().String("identity-validation-expression", "", "A validation expression for the incoming identity token.")
		apigateway_createAuthorizerCmd.Flags().String("name", "", "The name of the authorizer.")
		apigateway_createAuthorizerCmd.Flags().String("provider-arns", "", "A list of the Amazon Cognito user pool ARNs for the `COGNITO_USER_POOLS` authorizer.")
		apigateway_createAuthorizerCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_createAuthorizerCmd.Flags().String("type", "", "The authorizer type.")
		apigateway_createAuthorizerCmd.MarkFlagRequired("name")
		apigateway_createAuthorizerCmd.MarkFlagRequired("rest-api-id")
		apigateway_createAuthorizerCmd.MarkFlagRequired("type")
	})
	apigatewayCmd.AddCommand(apigateway_createAuthorizerCmd)
}
