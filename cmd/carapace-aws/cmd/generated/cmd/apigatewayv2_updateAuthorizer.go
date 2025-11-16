package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateAuthorizerCmd = &cobra.Command{
	Use:   "update-authorizer",
	Short: "Updates an Authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_updateAuthorizerCmd).Standalone()

		apigatewayv2_updateAuthorizerCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("authorizer-credentials-arn", "", "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("authorizer-id", "", "The authorizer identifier.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("authorizer-payload-format-version", "", "Specifies the format of the payload sent to an HTTP API Lambda authorizer.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("authorizer-result-ttl-in-seconds", "", "The time to live (TTL) for cached authorizer results, in seconds.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("authorizer-type", "", "The authorizer type.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("authorizer-uri", "", "The authorizer's Uniform Resource Identifier (URI).")
		apigatewayv2_updateAuthorizerCmd.Flags().String("enable-simple-responses", "", "Specifies whether a Lambda authorizer returns a response in a simple format.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("identity-source", "", "The identity source for which authorization is requested.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("identity-validation-expression", "", "This parameter is not used.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("jwt-configuration", "", "Represents the configuration of a JWT authorizer.")
		apigatewayv2_updateAuthorizerCmd.Flags().String("name", "", "The name of the authorizer.")
		apigatewayv2_updateAuthorizerCmd.MarkFlagRequired("api-id")
		apigatewayv2_updateAuthorizerCmd.MarkFlagRequired("authorizer-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateAuthorizerCmd)
}
