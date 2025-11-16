package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_reimportApiCmd = &cobra.Command{
	Use:   "reimport-api",
	Short: "Puts an Api resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_reimportApiCmd).Standalone()

	apigatewayv2_reimportApiCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_reimportApiCmd.Flags().String("basepath", "", "Specifies how to interpret the base path of the API during import.")
	apigatewayv2_reimportApiCmd.Flags().String("body", "", "The OpenAPI definition.")
	apigatewayv2_reimportApiCmd.Flags().String("fail-on-warnings", "", "Specifies whether to rollback the API creation when a warning is encountered.")
	apigatewayv2_reimportApiCmd.MarkFlagRequired("api-id")
	apigatewayv2_reimportApiCmd.MarkFlagRequired("body")
	apigatewayv2Cmd.AddCommand(apigatewayv2_reimportApiCmd)
}
