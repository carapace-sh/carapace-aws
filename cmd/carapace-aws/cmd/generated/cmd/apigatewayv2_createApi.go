package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createApiCmd = &cobra.Command{
	Use:   "create-api",
	Short: "Creates an Api resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createApiCmd).Standalone()

	apigatewayv2_createApiCmd.Flags().String("api-key-selection-expression", "", "An API key selection expression.")
	apigatewayv2_createApiCmd.Flags().String("cors-configuration", "", "A CORS configuration.")
	apigatewayv2_createApiCmd.Flags().String("credentials-arn", "", "This property is part of quick create.")
	apigatewayv2_createApiCmd.Flags().String("description", "", "The description of the API.")
	apigatewayv2_createApiCmd.Flags().String("disable-execute-api-endpoint", "", "Specifies whether clients can invoke your API by using the default execute-api endpoint.")
	apigatewayv2_createApiCmd.Flags().String("disable-schema-validation", "", "Avoid validating models when creating a deployment.")
	apigatewayv2_createApiCmd.Flags().String("ip-address-type", "", "The IP address types that can invoke the API.")
	apigatewayv2_createApiCmd.Flags().String("name", "", "The name of the API.")
	apigatewayv2_createApiCmd.Flags().String("protocol-type", "", "The API protocol.")
	apigatewayv2_createApiCmd.Flags().String("route-key", "", "This property is part of quick create.")
	apigatewayv2_createApiCmd.Flags().String("route-selection-expression", "", "The route selection expression for the API.")
	apigatewayv2_createApiCmd.Flags().String("tags", "", "The collection of tags.")
	apigatewayv2_createApiCmd.Flags().String("target", "", "This property is part of quick create.")
	apigatewayv2_createApiCmd.Flags().String("version", "", "A version identifier for the API.")
	apigatewayv2_createApiCmd.MarkFlagRequired("name")
	apigatewayv2_createApiCmd.MarkFlagRequired("protocol-type")
	apigatewayv2Cmd.AddCommand(apigatewayv2_createApiCmd)
}
