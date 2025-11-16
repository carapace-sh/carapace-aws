package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateApiCmd = &cobra.Command{
	Use:   "update-api",
	Short: "Updates an Api resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_updateApiCmd).Standalone()

		apigatewayv2_updateApiCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_updateApiCmd.Flags().String("api-key-selection-expression", "", "An API key selection expression.")
		apigatewayv2_updateApiCmd.Flags().String("cors-configuration", "", "A CORS configuration.")
		apigatewayv2_updateApiCmd.Flags().String("credentials-arn", "", "This property is part of quick create.")
		apigatewayv2_updateApiCmd.Flags().String("description", "", "The description of the API.")
		apigatewayv2_updateApiCmd.Flags().String("disable-execute-api-endpoint", "", "Specifies whether clients can invoke your API by using the default execute-api endpoint.")
		apigatewayv2_updateApiCmd.Flags().String("disable-schema-validation", "", "Avoid validating models when creating a deployment.")
		apigatewayv2_updateApiCmd.Flags().String("ip-address-type", "", "The IP address types that can invoke your API or domain name.")
		apigatewayv2_updateApiCmd.Flags().String("name", "", "The name of the API.")
		apigatewayv2_updateApiCmd.Flags().String("route-key", "", "This property is part of quick create.")
		apigatewayv2_updateApiCmd.Flags().String("route-selection-expression", "", "The route selection expression for the API.")
		apigatewayv2_updateApiCmd.Flags().String("target", "", "This property is part of quick create.")
		apigatewayv2_updateApiCmd.Flags().String("version", "", "A version identifier for the API.")
		apigatewayv2_updateApiCmd.MarkFlagRequired("api-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateApiCmd)
}
