package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createAuthorizerCmd = &cobra.Command{
	Use:   "create-authorizer",
	Short: "Creates an Authorizer for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createAuthorizerCmd).Standalone()

	apigatewayv2_createAuthorizerCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_createAuthorizerCmd.Flags().String("authorizer-credentials-arn", "", "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.")
	apigatewayv2_createAuthorizerCmd.Flags().String("authorizer-payload-format-version", "", "Specifies the format of the payload sent to an HTTP API Lambda authorizer.")
	apigatewayv2_createAuthorizerCmd.Flags().String("authorizer-result-ttl-in-seconds", "", "The time to live (TTL) for cached authorizer results, in seconds.")
	apigatewayv2_createAuthorizerCmd.Flags().String("authorizer-type", "", "The authorizer type.")
	apigatewayv2_createAuthorizerCmd.Flags().String("authorizer-uri", "", "The authorizer's Uniform Resource Identifier (URI).")
	apigatewayv2_createAuthorizerCmd.Flags().String("enable-simple-responses", "", "Specifies whether a Lambda authorizer returns a response in a simple format.")
	apigatewayv2_createAuthorizerCmd.Flags().String("identity-source", "", "The identity source for which authorization is requested.")
	apigatewayv2_createAuthorizerCmd.Flags().String("identity-validation-expression", "", "This parameter is not used.")
	apigatewayv2_createAuthorizerCmd.Flags().String("jwt-configuration", "", "Represents the configuration of a JWT authorizer.")
	apigatewayv2_createAuthorizerCmd.Flags().String("name", "", "The name of the authorizer.")
	apigatewayv2_createAuthorizerCmd.MarkFlagRequired("api-id")
	apigatewayv2_createAuthorizerCmd.MarkFlagRequired("authorizer-type")
	apigatewayv2_createAuthorizerCmd.MarkFlagRequired("identity-source")
	apigatewayv2_createAuthorizerCmd.MarkFlagRequired("name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_createAuthorizerCmd)
}
